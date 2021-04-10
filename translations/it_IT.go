package translations

func initItalianTranslation() Translation {
	translation := createTranslation()

	translation.put("requires-js", "This website requires JavaScript to run properly.")

	translation.put("start-the-game", "Inizia il gioco")
	translation.put("start", "Inizia")
	translation.put("game-not-started-title", "La partita non è ancora iniziata")
	translation.put("waiting-for-host-to-start", "Aspetto che venga avviata la partita")

	translation.put("round", "Round")
	translation.put("toggle-soundeffects", "Toggle soundeffects")
	translation.put("change-your-name", "Cambia nome")
	translation.put("randomize", "Random")
	translation.put("change", "Cambia")
	translation.put("votekick-a-player", "Vota per buttare fuori")
	translation.put("time-left", "tempo")

	translation.put("last-turn", "(Ultimo turno: %s)")

	translation.put("drawer-kicked", "Since the kicked player has been drawing, none of you will get any points this round.")
	translation.put("self-kicked", "Ciao, t'hanno buttato fuori!")
	translation.put("kick-vote", "(%s/%s) giocatori hanno votato per buttare fuori %s.")
	translation.put("player-kicked", "Il giocatore è stato buttato fuori.")
	translation.put("owner-change", "%s è il nuovo capo supremo.")

	translation.put("change-lobby-settings", "Configurazione")
	translation.put("lobby-settings-changed", "Configurazione modificata")
	translation.put("advanced-settings", "Configurazione avanzata")
	translation.put("word-language", "Lingua")
	translation.put("drawing-time-setting", "Tempo per disegnare")
	translation.put("rounds-setting", "Rounds")
	translation.put("max-players-setting", "Max giocatori")
	translation.put("public-lobby-setting", "Pubblica")
	translation.put("custom-words", "Parole personalizzate")
	translation.put("custom-words-info", "Parole, separate da virgole")
	translation.put("custom-words-chance-setting", "Probabilità delle parole custom")
	translation.put("players-per-ip-limit-setting", "Players per IP Limit")
	translation.put("enable-votekick-setting", "Allow Votekick")
	translation.put("save-settings", "Salva")
	translation.put("input-contains-invalid-data", "Your input contains invalid data:")
	translation.put("please-fix-invalid-input", "Correct the invalid input and try again.")
	translation.put("create-lobby", "Crea partita")

	translation.put("players", "Giocatori")
	translation.put("refresh", "Ricarica")
	translation.put("join-lobby", "Unisciti")

	translation.put("message-input-placeholder", "Scrivi qui i tuoi tentativi (anche se tanto vince il Fusi)")

	translation.put("choose-a-word", "Scegli")
	translation.put("waiting-for-word-selection", "In attesa...")
	//This one doesn't use %s, since we want to make one part bold.
	translation.put("is-choosing-word", "sta scegliendo la parola.")

	translation.put("close-guess", "Fuochino!")
	translation.put("correct-guess", "Grande! Sei il Fusi?")
	translation.put("correct-guess-other-player", "%s ha indovinato!")
	translation.put("round-over", "Turn over, no word was chosen.")
	translation.put("round-over-no-word", "Fine, la parola era: '%s'.")
	translation.put("game-over-win", "Congratulazione, hai vinto, ma ha comunque vinto il Fusi!")
	translation.put("game-over", "Sei arrivato %s. con %s punti")

	translation.put("change-active-color", "Change your active color")
	translation.put("use-pencil", "Usa ì lapis")
	translation.put("use-eraser", "Usa la gomma")
	translation.put("use-fill-bucket", "Use fill bucket (Fills the target area with the selected color)")
	translation.put("change-pencil-size-to", "Cambia dimensione a %s")
	translation.put("clear-canvas", "Pulisci tutto")

	translation.put("connection-lost", "Connection lost!")
	translation.put("connection-lost-text", "Attempting to reconnect"+
		" ...\n\nMake sure your internet connection works.\nIf the "+
		"problem persists, contact the webmaster.")
	translation.put("error-connecting", "Error connecting to server")
	translation.put("error-connecting-text",
		"Scribble.rs couldn't establish a socket connection.\n\nWhile your internet "+
			"connection seems to be working, either the\nserver or your firewall hasn't "+
			"been configured correctly.\n\nTo retry, reload the page.")

	//Generic words
	//As "close" in "closing the window"
	translation.put("close", "Chiudi")
	translation.put("no", "No")
	translation.put("yes", "Sì")
	translation.put("system", "sistema")

	translation.put("source-code", "Source Code")
	translation.put("help", "Aiuto")
	translation.put("contact", "Contatti")
	translation.put("submit-feedback", "Feedback einreichen")
	translation.put("stats", "Status")

	RegisterTranslation("en", translation)
	RegisterTranslation("en-de", translation)

	return translation
}
