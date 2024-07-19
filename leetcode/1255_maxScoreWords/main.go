package main

func maxScoreWords(words []string, letters []byte, score []int) int {
    letter_score := make(map[byte]int)
    max := 0
    for index, item := range score{
        if (item == 0){
            continue
        } else {
            letter_score[letters[index]] = item 
        }
    }

    for _, word := range words{
        count := 0
        for i := 0; i < len(word); i++ {
            count = count + letter_score[word[i]]
        }
        if max < count{
            max = count
        }
    }
    return max
}