package comments

func codifyText(text string) string {
	return "```\n" + text + "\n```"
}

func attachComment(text string, comment string) string {
	return comment + "\n" + text
}

func makeComment(text string, comment string, codify bool) string {
	if codify == true {
		return attachComment(codifyText(text), comment)
	}
	return attachComment(text, comment)
}

func MakeSingleComment(rawText []string, rawComment []string, codify bool) string {
	if len(rawText) == 1 {
		return makeComment(rawText[0], rawComment[0], codify)
	}
	var result string
	for i, c := range rawText {
		if len(rawComment) <= i {
			result += makeComment(c, rawComment[i], codify)
		}
		result += makeComment(c, "", codify)
	}
	return result
}
