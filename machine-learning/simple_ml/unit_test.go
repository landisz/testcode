package main

import (
	//"context"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
	"testing"
)

type sample struct {
	X float64
	Y float64
	Result float64
}

const resultUnit = 1

func TestGuess(t *testing.T) {

	//sample of your secret linear algorithm. e.g. Z=a*X+b*Y+c
	trainingSet := []sample{
		{X: 2, Y: 5, Result: 11},
		{X: 7, Y: 4, Result: 14},
		{X: 6, Y: 11, Result: 27},
		{X: 13, Y: 1, Result: 15},
	}
	guess := []float64{4,6}

	nwk := Neural_network{GetWeight()}

	trainingInput, trainingExpect := GetTrainData(trainingSet)
	nwk.train(trainingInput, trainingExpect, int(TRAIN_ROUNDS))
	fmt.Println("My guess is: ", nwk.Guess(guess))
}

func GetTrainData(data []sample) (*mat.Dense, *mat.Dense) {
	sets := len(data)
	var  inputStream, outputStream []float64
	for _, tset := range data{
		inputStream = append(inputStream, tset.X)
		inputStream = append(inputStream, tset.Y)
		outputStream = append(outputStream, tset.Result)
	}
	return mat.NewDense(sets,2, inputStream), mat.NewDense(sets,1, outputStream)
}

func (n *Neural_network)Guess(test []float64) float64 {
	predict := n.think(mat.NewDense(1, 2, test))

	return math.Round(predict.At(0,0)/resultUnit)*resultUnit
}