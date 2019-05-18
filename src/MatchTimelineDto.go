package riotapi

//MatchTimelineDto represents matchTimeline per min
type MatchTimelineDto struct {
	Frames []struct {
		Timestamp         int `json:"timestamp"`
		ParticipantFrames struct {
			P1 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"1"`
			P2 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"2"`
			P3 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"3"`
			P4 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"4"`
			P5 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"5"`
			P6 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"6"`
			P7 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"7"`
			P8 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"8"`
			P9 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"9"`
			P10 struct {
				TotalGold     int `json:"totalGold"`
				TeamScore     int `json:"teamScore"`
				ParticipantID int `json:"participantId"`
				Level         int `json:"level"`
				CurrentGold   int `json:"currentGold"`
				MinionsKilled int `json:"minionsKilled"`
				DominionScore int `json:"dominionScore"`
				Position      struct {
					Y int `json:"y"`
					X int `json:"x"`
				} `json:"position"`
				Xp                  int `json:"xp"`
				JungleMinionsKilled int `json:"jungleMinionsKilled"`
			} `json:"10"`
		} `json:"participantFrames"`
		Events []interface{} `json:"events"`
	} `json:"frames"`
	FrameInterval int `json:"frameInterval"`
}