package main

import "os/exec"

func suara(text string) {

	command :=
		"Add-Type -AssemblyName System.Speech; " +
			"$voice = New-Object System.Speech.Synthesis.SpeechSynthesizer; " +
			"$voice.SelectVoice('Microsoft Zira Desktop'); " +
			"$voice.Speak('" + text + "')"

	exec.Command(
		"PowerShell",
		"-Command",
		command,
	).Run()
}
