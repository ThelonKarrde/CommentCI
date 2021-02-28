package format

func CodifyText(text string) string {
	return "```\n" + text + "\n```\n"
}

func appendComment(text string, comment string) string {
	return comment + "\n" + text + "\n"
}

func Comment(text string, comment string, codify bool) string {
	if codify == true {
		return appendComment(CodifyText(text), comment)
	}
	return appendComment(text, comment)
}

func SingleComment(rawText []string, rawComment []string, codify bool) string {
	if len(rawText) == 1 {
		if len(rawComment) > 0 {
			return Comment(rawText[0], rawComment[0], codify)
		}
		return Comment(rawText[0], "", codify)
	}
	var result string
	for i, t := range rawText {
		if i >= len(rawComment) {
			result += Comment(t, "", codify)
		} else {
			result += Comment(t, rawComment[i], codify)
		}
	}
	return result
}
