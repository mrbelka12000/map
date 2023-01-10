package visualiser

import (
	"bufio"
	"container/list"
	"fmt"
	"github.com/graincomg/graincom_logistik/models"
	"github.com/graincomg/graincom_logistik/prettifier"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"log"
	"os"
	"strconv"
)

func Visualiser(nodes []models.Node, ways []models.Way, l *list.List) {
	_, err := os.Create("./visualiser/data.txt")
	if err != nil {
		log.Fatalf("Error")
	}
	//for _, child := range nodes {
	//	str := child.Lat + "," + child.Lon + "\n"
	//	if _, err := file.Write([]byte(str)); err != nil {
	//		log.Fatal(err)
	//	}
	//}

	file3, err := os.Create("./visualiser/map.txt")
	if err != nil {
		log.Fatalf("Error")
	}
	for _, k := range nodes {
		str := k.Lat + "," + k.Lon + "\n"
		break
		if _, err := file3.Write([]byte(str)); err != nil {
			log.Fatal(err)
		}
	}

	xys3, err := readData("./visualiser/map.txt")
	if err != nil {
		log.Fatalf("could not read map.txt: %v", err)
	}

	err = plotData("./out.png", xys3, nil, nil, nodes, ways)
	if err != nil {
		log.Fatalf("could not plot data: %v", err)
	}
}

type xy struct{ x, y float64 }

func readData(path string) (plotter.XYs, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var xys plotter.XYs
	s := bufio.NewScanner(f)
	for s.Scan() {
		var x, y float64
		_, err := fmt.Sscanf(s.Text(), "%f,%f", &x, &y)
		if err != nil {
			log.Printf("discarding bad data point %q: %v", s.Text(), err)
			continue
		}
		xys = append(xys, struct{ X, Y float64 }{x, y})
	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("could not scan: %v", err)
	}
	return xys, nil
}

func plotData(path string, xys plotter.XYs, xys2 plotter.XYs, xys3 plotter.XYs, moves []models.Node, ways []models.Way) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create %s: %v", path, err)
	}

	p := plot.New()

	count := 0
	for i := 0; i < len(ways)-1; i++ {

		for j := 0; j < len(ways[i].Nd)-1; j++ {
			if count%2 != 0 {
				//continue
			}
			node1, _ := prettifier.GetNodeByID(ways[i].Nd[j].Ref)
			node2, _ := prettifier.GetNodeByID(ways[i].Nd[j+1].Ref)
			count = 0
			x1, _ := strconv.ParseFloat(node1.Lat, 64)
			y1, _ := strconv.ParseFloat(node1.Lon, 64)
			x2, _ := strconv.ParseFloat(node2.Lat, 64)
			y2, _ := strconv.ParseFloat(node2.Lon, 64)
			l, err := plotter.NewLine(plotter.XYs{
				{x1, y1}, {x2, y2},
			})

			if err != nil {
				return fmt.Errorf("could not create line: %v", err)
			}
			l.Width = 0.7
			p.Add(l)
		}
		count++
	}
	wt, err := p.WriterTo(512, 512, "png")
	if err != nil {
		return fmt.Errorf("could not create writer: %v", err)
	}
	_, err = wt.WriteTo(f)
	if err != nil {
		return fmt.Errorf("could not write to %s: %v", path, err)
	}

	return nil
}
