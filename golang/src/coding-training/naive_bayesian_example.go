package main

import "fmt"

func main() {
	xTrain := [][]int{
		{0, 1, 0, 1}, // 0
		{1, 0, 1, 1}, // 1
		{0, 0, 0, 1}, // 0
		{1, 0, 1, 0}, // 1
		{1, 1, 0, 0}, // 1
	}
	yTrain := []int{0, 1, 0, 1, 1}
	xTest := []int{1, 1, 1, 0}

	result := predicWithNaiveBasianClassifier(xTrain, yTrain, xTest)
	fmt.Println(result)
}

/**
 * P(A|C)와 P(B|C)의 확률값을 비교해서 더 큰 값의 레이블을 반환한다.

 P(A|array) = (P(array|A) * P(A)) / P(array)
 P(B|array) = (P(array|B) * P(B)) / P(array)

 P(A|array) = P(array|A) * P(A)
 P(B|array) = P(array|B) * P(B)

 P(A) = 사건 A가 일어날 확률. 여기서는 y값이 0이 될 확률인데 2/5
 P(B) = 사건 B가 일어날 확률. 여기서는 y값이 1이 될 확률인데 3/5
 P(array|A) = y가 0일 때, 지정한 array가 나올 확률
 P(array|B) = y가 1일 때, 지정한 array가 나올 확률

 위의 예제에서는 array가 {1, 1, 1, 0}으로 주어졌기 때문에 P({1, 1, 1, 0}|0)와 P({1, 1, 1, 0}|1)을 구하면 된다.
 그러고 나서 높은 쪽의 레이블을 선택하면 된다.

 [클래스0의 경우]
 feature 1에 1이 나오면서 y가 0인 횟수 = 0, 확률은 0/5
 feature 2에 1이 나오면서 y가 0인 횟수 = 1, 확률은 1/5
 feature 3에 1이 나오면서 y가 0인 횟수 = 0, 확률은 0/5
 feature 4에 0이 나오면서 y가 0인 횟수 = 0, 확률은 0/5

 [클래스1의 경우]
 feature 1에 1이 나오면서 y가 1인 횟수 = 3, 확률은 3/5
 feature 2에 1이 나오면서 y가 1인 횟수 = 1, 확률은 1/5
 feature 3에 1이 나오면서 y가 1인 횟수 = 2, 확률은 2/5
 feature 4에 0이 나오면서 y가 1인 횟수 = 2, 확률은 2/5

 위 공식에 대입하면
 P(0|{1,1,1,0}) = (0/5 * 1/5 * 0/5 * 0/5) * 2/5
 P(1|{1,1,1,0}) = (2/5 * 1/5 * 2/5 * 1/5) * 3/5

 고로 {1,1,1,0} 일 때는 1로 예측

 */
func predicWithNaiveBasianClassifier(xTrain [][]int, yTrain []int, xTest []int) int {
	rowCount := len(xTrain)
	featureValCountForClass0 := [4]int{0, 0, 0, 0}
	featureValCountForClass1 := [4]int{0, 0, 0, 0}

	for i := 0; i < rowCount; i++ {
		featureCount := len(xTrain[i])

		for j := 0; j < featureCount; j++ {
			if xTrain[i][j] == xTest[j] && yTrain[i] == 0 {
				featureValCountForClass0[j] += 1
			} else if xTrain[i][j] == xTest[j] && yTrain[i] == 1 {
				featureValCountForClass1[j] += 1
			}
		}
	}
	probForClass0 := [4]float64{
		float64(featureValCountForClass0[0]) / float64(rowCount),
		float64(featureValCountForClass0[1]) / float64(rowCount),
		float64(featureValCountForClass0[2]) / float64(rowCount),
		float64(featureValCountForClass0[3]) / float64(rowCount),
	}
	probForClass1 := [4]float64{
		float64(featureValCountForClass1[0]) / float64(rowCount),
		float64(featureValCountForClass1[1]) / float64(rowCount),
		float64(featureValCountForClass1[2]) / float64(rowCount),
		float64(featureValCountForClass1[3]) / float64(rowCount),
	}

	resultForClass0 := float64(1)
	resultForClass1 := float64(1)
	pc0 := 1
	pc1 := 1
	for i := 0; i < len(probForClass0); i++ {
		resultForClass0 *= probForClass0[i]
		resultForClass1 *= probForClass1[i]
	}
	for i := 0; i < len(yTrain); i++ {
		if yTrain[i] == 0 {
			pc0 += 1
		} else {
			pc1 += 1
		}
	}
	p0 := float64(pc0) / float64(len(yTrain))
	p1 := float64(pc1) / float64(len(yTrain))

	if resultForClass0*p0 > resultForClass1*p1 {
		return 0
	}
	return 1
}
