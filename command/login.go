		// Store the token in the local client
		if err := tokenHelper.Store(token); err != nil {
			c.UI.Error(fmt.Sprintf("Error storing token: %s", err))
			if !c.flagNoPrint {
				c.UI.Error(wrapAtLength(
					"Authentication was successful, but the token was not persisted. The "+
						"resulting token is shown below for your records.") + "\n")
				OutputSecret(c.UI, secret)
			}
			return 2
		}