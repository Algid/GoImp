func getToken() {

	var {
		ch = " "
		lexeme = ""
	}
	
	//Keep going until you find a nonwhitespace character
	for Isspace(ch) {
		ch = getChar();
	}
	
	//Check for the error character
	if ch == "!" {
		lexeme += ch
		ch = getChar()
		return t_error
	}

	//If it is a /, you need to check and see if its a comment
	if ch == "/" {
		ch = getChar()
		//If its two // in a row, then it is a comment
		//and the rest of the line can be ignored
		if ch == "/" {
			for ch != "/n" {
				ch = getChar()
			}
			return getToken()
		}
		//Otherwise its just the division sign
		else {
			lexeme += "/"
			return t_div
		}
	}
	
	//If the character is a " then it is a string
	//Get all characters until next ", then return
	if ch == """ {
		ch = getChar()
		for ch != """ {
			lexeme += ch
			ch = getChar()
		}
		ch = getChar()
		return t_string
	}

	//If its a number, then get all digits and return
	if IsDigit(ch) {
		for IsDigit(ch) {
			lexeme += ch
			ch = getChar()
		}
		return t_number;
	}

	//If its a letter, then it is either a variable name
	//or a reserved word
	if Isletter(ch) || ch == '_' {
		//get all values, then check if its in the map
		for (Isletter(ch) || Isdigit(ch) || ch == '_' {
			lexeme += ch
			ch = getChar()
		}
		if //is in map {
			return tokenMapValue[lexeme]
		}
		else {
			return t_id
		}
	}
	
	//If its a punctuation, determine which one it
	//is and return the corresponding token
	if Ispunct(ch) {
		switch ch {
		case "<": {
			ch = getChar()
			if ch == "=" {
				ch = getChar()
				return t_le
			}
			else
				return t_lt
		}
		case ">": {
			ch = getChar()
			if ch == "=" {
				ch = getChar()
				return t_ge
			}
			else
				return t_gt
		}
		case "=": {
			ch = getChar()
			if ch == "=" {
				ch = getChar()
				return t_eq
			}
			else
				return t_assign
		}
		case "#": {
			ch = getChar()
			return t_ne
		}
		case "," {
			ch = getChar()
			return t_comma
		}
		case ":" {
			ch = getChar()
			return t_colon
		}
		case "(" {
			ch = getChar()
			return t_rparen
		}
		case ")" {
			ch = getChar()
			return t_lparen
		}
		case ";" {
			ch = getChar()
			return t_semi
		}
		case "+" {
			ch = getChar()
			return t_plus
		}
		case "-" {
			ch = getChar()
			return t_minus
		}
		case "*" {
			ch = getChar()
			return t_mult
		}
		case "%" {
			ch = getChar()
			return t_mod
		}
		default:
		}
		ch = getChar()
	}
	return t_eof
}
