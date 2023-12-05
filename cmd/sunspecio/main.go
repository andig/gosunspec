// This package contains a command that allows an input address space to be copied into a
// similarly shaped output address space.
//
// For example:
//
//	sunspecio --in:type=modbus --in:device=/dev/ttyS0 --in:speed=38400 --out:type=xml
//
// will read the Sunspec address space in the Modbus device connected to the specified serial port
// and copy it into an xml document which is written to stdout.
package main

import (
	xmlencoding "encoding/xml"
	"flag"
	"io"
	"log"
	"os"
	"strings"
	"time"

	sunspec "github.com/andig/gosunspec"
	"github.com/andig/gosunspec/layout"
	"github.com/andig/gosunspec/modbus"
	_ "github.com/andig/gosunspec/models"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/xml"
	modbusapi "github.com/grid-x/modbus"
)

func openModbus(device string, config func(handler *modbusapi.RTUClientHandler)) (modbusapi.Client, func(), error) {
	handler := modbusapi.NewRTUClientHandler(device)

	handler.BaudRate = 38400
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveID = 1
	handler.Timeout = 5 * time.Second

	config(handler)

	err := handler.Connect()
	if err != nil {
		return nil, nil, err
	}

	client := modbusapi.NewClient(handler)

	return client, func() { handler.Close() }, nil
}

func encode(w io.Writer, i interface{}) error {
	encoder := xmlencoding.NewEncoder(w)
	encoder.Indent("", "   ")
	err := encoder.Encode(i)
	return err
}

func readLayout(n string) layout.AddressSpaceLayout {
	if n != "" {
		if f, err := os.Open(n); err != nil {
			log.Fatal(err)
		} else {
			defer f.Close()
			obj := &layout.RawLayout{}
			if err := xmlencoding.NewDecoder(f).Decode(obj); err != nil {
				log.Fatal(err)
			}
			return obj
		}
	}
	return nil
}

func loadModels(smdxDir string) {
	if !strings.HasSuffix(smdxDir, "/") {
		smdxDir = smdxDir + "/"
	}
	files, err := os.ReadDir(smdxDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), "smdx_") {
			smdxFile, err := os.OpenFile(smdxDir+file.Name(), os.O_RDONLY, 0o644)
			if err != nil {
				log.Fatal(err)
			}

			def, err := smdx.FromXML(smdxFile)
			if err != nil {
				smdxFile.Close()
				log.Fatal(err)
			}
			smdxFile.Close()

			if len(def.Models) < 1 {
				log.Printf("failed to find any models in %s", file.Name())
				continue
			} else {
				for i := range def.Models {
					smdx.RegisterModel(&def.Models[i])
				}
			}
		}
	}
}

func main() {
	var inType string
	var inDevice string
	var inSpeed int
	var inLayout string
	var inSlave int

	var outType string
	var outDevice string
	var outSpeed int
	var outLayout string
	var outSlave int

	var modelsDir string

	flag.StringVar(&inType, "in:type", "xml", "The type of the input address space.")
	flag.StringVar(&inDevice, "in:device", "", "The device from which the copied address space is read (in:type=modbus-rtu, only).")
	flag.IntVar(&inSpeed, "in:speed", 38400, "The connection speed for the device specified by in:device.")
	flag.StringVar(&inLayout, "in:layout", "", "The address layout to be used with the input address space. Empty for the default sunspec layout. Not required for an xml address spaces.")
	flag.IntVar(&inSlave, "in:slave", 1, "The slave identifier")

	flag.StringVar(&outType, "out:type", "xml", "The type of the output address space.")
	flag.StringVar(&outDevice, "out:device", "", "The device to which the copied address space is written (out:type=modbus-rtu, only).")
	flag.IntVar(&outSpeed, "out:speed", 38400, "The connection speed for the write device specified by out:device.")
	flag.StringVar(&outLayout, "out:layout", "", "The address layout to be used with the output address space. Empty for the default sunspec layout. Not required for an xml address spaces.")
	flag.IntVar(&outSlave, "out:slave", 1, "The slave identifier")

	flag.StringVar(&modelsDir, "models-dir", "", "The location from which auxiliary models are loaded.")

	flag.Parse()

	var in sunspec.Array
	var err error

	if modelsDir != "" {
		loadModels(modelsDir)
	}

	if inType == "xml" {
		decoder := xmlencoding.NewDecoder(os.Stdin)
		elements := &xml.DataElement{}
		err = decoder.Decode(elements)
		if err != nil {
			log.Fatalf("failed to decode: %v", err)
		}
		in, err = xml.Open(elements)
	} else if inType == "modbus-rtu" {
		inLayoutObj := readLayout(inLayout)

		var client modbusapi.Client
		var closer func()
		client, closer, err = openModbus(inDevice, func(handler *modbusapi.RTUClientHandler) {
			handler.BaudRate = inSpeed
			handler.SlaveID = byte(inSlave)
		})
		if err != nil {
			log.Fatal(err)
		}
		defer closer()
		if inLayoutObj == nil {
			in, err = modbus.Open(client)
		} else {
			in, err = modbus.OpenWithLayout(client, inLayoutObj)
		}
	} else {
		log.Fatalf("invalid input type: %s", inType)
	}

	if err != nil {
		log.Fatal(err)
	}

	in.Do(func(d sunspec.Device) {
		d.Do(func(m sunspec.Model) {
			m.Do(func(b sunspec.Block) {
				err = b.Read()
				if err == nil || inType == "xml" {
					// errors expected on XML reads.
					return
				} else {
					log.Fatal(err)
				}
			})
		})
	})

	if outType == "xml" {
		_, outX := xml.CopyArray(in)
		encode(os.Stdout, outX)
	} else if outType == "modbus-rtu" {
		outLayoutObj := readLayout(outLayout)
		var client modbusapi.Client
		var closer func()
		client, closer, err = openModbus(outDevice, func(handler *modbusapi.RTUClientHandler) {
			handler.BaudRate = outSpeed
			handler.SlaveID = byte(outSlave)
		})
		if err != nil {
			log.Fatal(err)
		}
		defer closer()
		var out sunspec.Array
		if outLayoutObj == nil {
			out, err = modbus.Open(client)
		} else {
			out, err = modbus.OpenWithLayout(client, outLayoutObj)
		}
		if err != nil {
			log.Fatal(err)
		}
		inDevices := in.Collect(sunspec.AllDevices)
		outDevices := out.Collect(sunspec.AllDevices)

		for i, inDevice := range inDevices {
			if i >= len(outDevices) {
				break
			}
			outDevice := outDevices[i]
			inDevice.Do(func(inModel sunspec.Model) {
				outModel := outDevice.MustModel(inModel.Id())
				bn := 0
				inModel.Do(func(b sunspec.Block) {
					outBlock := outModel.MustBlock(bn)
					ids := []string{}
					b.DoScaleFactorsFirst(func(p sunspec.Point) {
						if p.Error() == nil {
							outBlock.MustPoint(p.Id()).SetValue(p.Value())
							ids = append(ids, p.Id())
						}
					})
					b.Write(ids...)
					bn++
				})
			})
		}
	}
}
