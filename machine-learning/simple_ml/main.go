package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	//"gonum.org/v1/gonum/blas/blas64"
	//"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

const TRAIN_ROUNDS=50000

type Neural_network struct {
	Weights *mat.Dense
}

//type twoDarray [][]float64


func main () {
	rand.Seed(time.Now().UTC().UnixNano())

	nwk := Neural_network{GetWeight()}

	// hidden linear algorithm: z= 2x-y
	trainInput := mat.NewDense(4, 2, []float64{
		1,2,
		3,4,
		7,5,
		8,6,
	})
	//trainExpect := transpose([][]float64{{10, 4, 14, 30}})
	trainExpect := mat.NewDense(4, 1, []float64{
		0,2,9,10,
	})

	// Training the neural network using the training set.
	nwk.train(trainInput, trainExpect, int(TRAIN_ROUNDS))

	// Ask the neural network the output
	//var predict mat.Matrix
	predict := nwk.think(mat.NewDense(1, 2, []float64{2,5}))
	fmt.Println("Result at ", TRAIN_ROUNDS, "training round: ")
	fmt.Printf("%v\n",mat.Formatted(predict,mat.Prefix(""), mat.Squeeze()))
	//val := int(predict.At(0,0))
	
	fmt.Println("Predicted value is: ", math.Round(predict.At(0,0)/0.01)*0.01)
}

//func transpose(slice [][]float64) [][]float64 {
//	xl := len(slice[0]) //colums
//	yl := len(slice) // rows
//	result := make([][]float64, xl)
//	for i := range result {
//		result[i] = make([]float64, yl)
//	}
//	for i := 0; i < xl; i++ {
//		for j := 0; j < yl; j++ {
//			result[i][j] = slice[j][i]
//		}
//	}
//	return result
//}

func (nn *Neural_network) train(inputs *mat.Dense, expects *mat.Dense, num int){

	var err, adjustment mat.Dense
	var outputs *mat.Dense
	for i:=0; i<num; i++ {
		outputs = nn.think(inputs)
		//fmt.Println("Formatted think reply: \n", mat.Formatted(outputs))

		err.Sub(expects, outputs)
		adjustment.Mul(inputs.T(), &err)
		adjustment.Scale(0.001, &adjustment)

		nn.Weights.Add(nn.Weights, &adjustment)
	}
}


func (nn *Neural_network) think(inputs *mat.Dense) *mat.Dense{
	//xl := len(inputs[0])
	//yl := len(inputs)
	//result := [][]float64{}
	//for i := 0; i < xl; i++ {
	//	for j := 0; j < yl; j++ {
	//		result[i][j] = inputs[i][j]*nn.Weights
	//	}
	//}
	var result mat.Dense
	//mat.Formatted(result.Mul(inputs, nn.Weights))
	//eturn mat.NewDense(xl, 1, *[]float64(result.Mul(inputs, nn.Weights)))
	result.Mul(inputs, nn.Weights)
	//fmt.Println("Formatted think result: \n", mat.Formatted(&result))

	return &result
}


//func GetRandomMatrix(row int, column int) *mat.Dense{
//
//	//matrix := make([][]float64, row)
//	//for m := range matrix {
//	//	matrix[m] = make([]float64, column)
//	//}
//	data := make([]float64, (row*column))
//
//	//for i:=0; i<row; i++ {
//		for j:=0; j<(row*column); j++ {
//			//s := rand.NewSource(time.Now().UnixNano())
//			//r := rand.New(s)
//
//			data[j] = rand.Float64()
//		}
//	//}
//	return mat.NewDense(row, column, data)
//}

func GetWeight() *mat.Dense{
	dataw := make([]float64, 2)
	for i := range dataw {
		dataw[i] = 2*rand.NormFloat64()-1
	}

	return mat.NewDense(2, 1, dataw)
}