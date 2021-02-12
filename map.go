package gmoji

/**
	Font Awesome Emojis
	Versions
	{
		from 1.0 to 3.0 is available
	}
**/

var gmojiMap = map[string]string{
	":Glass:":                   "\uf000",
	":Music:":                   "\uf001",
	":Search:":                  "\uf002",
	":EnvelopeOutlined:":        "\uf003",
	":Heart:":                   "\uf004",
	":Star:":                    "\uf005",
	":StarOutlined:":            "\uf006",
	":User:":                    "\uf007",
	":Film:":                    "\uf008",
	":ThLarge:":                 "\uf009",
	":Th:":                      "\uf00a",
	":ThList:":                  "\uf00b",
	":Check:":                   "\uf00c",
	":Times:":                   "\uf00d",
	":SearchPlus:":              "\uf00e",
	":SearchMnus:":              "\uf010",
	":PowerOff:":                "\uf011",
	":Signal:":                  "\uf012",
	":Cog:":                     "\uf013",
	":TrashOutlined:":           "\uf014",
	":Home:":                    "\uf015",
	":FileOutlined:":            "\uf016",
	":ClockOutlined:":           "\uf017",
	":Road:":                    "\uf018",
	":Download:":                "\uf019",
	":ArrowCircleOutlinedDown:": "\uf01a",
	":ArrowCircleOutlinedUp:":   "\uf01b",
	":Inbox:":                   "\uf01c",
	":PlayCircleOutlined:":      "\uf01d",
	":Repeat:":                  "\uf01e",
	":Refresh:":                 "\uf021",
	":ListAlt:":                 "\uf022",
	":Lock:":                    "\uf023",
	":Flag:":                    "\uf024",
	":Headphones:":              "\uf025",
	":VolumeOff:":               "\uf026",
	":VolumeDown:":              "\uf027",
	":VolumeUp:":                "\uf028",
	":Qrcode:":                  "\uf029",
	":Barcode:":                 "\uf02a",
	":Tag:":                     "\uf02b",
	":Tags:":                    "\uf02c",
	":Book:":                    "\uf02d",
	":Bookmark:":                "\uf02e",
	":PrintEmoji:":              "\uf02f",
	":Camera:":                  "\uf030",
	":Font:":                    "\uf031",
	":Bold:":                    "\uf032",
	":Italic:":                  "\uf033",
	":TextHeight:":              "\uf034",
	":TextWidth:":               "\uf035",
	":AlignLeft:":               "\uf036",
	":AlignCenter:":             "\uf037",
	":AlignRight:":              "\uf038",
	":AlignJustify:":            "\uf039",
	":List:":                    "\uf03a",
	":Outdent:":                 "\uf03b",
	":Indent:":                  "\uf03c",
	":VideoCamera:":             "\uf03d",
	":Picture:":                 "\uf03e",
	":Pencil:":                  "\uf040",
	":MapMarker:":               "\uf041",
	":Adjust:":                  "\uf042",
	":Tint:":                    "\uf043",
	":PencilSquareOutlined:":    "\uf044",
	":ShareSquareOutlined:":     "\uf045",
	":CheckSquareOutlined:":     "\uf046",
	":Arrows:":                  "\uf047",
	":StepBackward:":            "\uf048",
	":FastBackward:":            "\uf049",
	":Backward:":                "\uf04a",
	":Play:":                    "\uf04b",
	":Pause:":                   "\uf04c",
	":Stop:":                    "\uf04d",
	":Forward:":                 "\uf04e",
	":FastForward:":             "\uf050",
	":StepForward:":             "\uf051",
	":Eject:":                   "\uf052",
	":ChevronLeft:":             "\uf053",
	":ChevronRight:":            "\uf054",
	":PlusCircle:":              "\uf055",
	":MinusCircle:":             "\uf056",
	":TimesCircle:":             "\uf057",
	":CheckCircle:":             "\uf058",
	":QuestionCircle:":          "\uf059",
	":InfoCircle:":              "\uf05a",
	":TimesCircleOutlined:":     "\uf05c",
	":CheckCircleOutlined:":     "\uf05d",
	":Ban:":                     "\uf05e",
	":ArrowLeft:":               "\uf060",
	":ArrowRight:":              "\uf061",
	":ArrowUp:":                 "\uf062",
	":ArrowDown:":               "\uf063",
	":Share:":                   "\uf064",
	":Expand:":                  "\uf065",
	":Compress:":                "\uf066",
	":Plus:":                    "\uf067",
	":Minus:":                   "\uf068",
	":Asterisk:":                "\uf069",
	":ExclamationCircle:":       "\uf06a",
	":Gift:":                    "\uf06b",
	":Leaf:":                    "\uf06c",
	":Fire:":                    "\uf06d",
	":Eye:":                     "\uf06e",
	":EyeSlash:":                "\uf070",
	":ExclamationTriangle:":     "\uf071",
	":Plane:":                   "\uf072",
	":Calendar:":                "\uf073",
	":Random:":                  "\uf074",
	":Comment:":                 "\uf075",
	":Magnet:":                  "\uf076",
	":ChevronUp:":               "\uf077",
	":ChevronDown:":             "\uf078",
	":Retweet:":                 "\uf079",
	":ShoppingCart:":            "\uf07a",
	":Folder:":                  "\uf07b",
	":FolderOpen:":              "\uf07c",
	":ArrowsVertical:":          "\uf07d",
	":ArrowsHorizontal:":        "\uf07e",
	":BarChart:":                "\uf080",
	":TwitterSquare:":           "\uf081",
	":FacebookSquare:":          "\uf082",
	":CameraRetro:":             "\uf083",
	":Key:":                     "\uf084",
	":Cogs:":                    "\uf085",
	":Comments:":                "\uf086",
	":ThumbsUpOutlined:":        "\uf087",
	":ThumbsDownOutlined:":      "\uf088",
	":StarHalf:":                "\uf089",
	":HeartOutlined:":           "\uf08a",
	":SignOut:":                 "\uf08b",
	":LinkedInSquare:":          "\uf08c",
	":ThumbTack:":               "\uf08d",
	":ExternalLink:":            "\uf08e",
	":SignIn:":                  "\uf090",
	":Trophy:":                  "\uf091",
	":GitHubSquare:":            "\uf092",
	":Upload:":                  "\uf093",
	":LemonOutlined:":           "\uf094",
	":Phone:":                   "\uf095",
	":SquareOutlined:":          "\uf096",
	":BookmarkOutlined:":        "\uf097",
	":PhoneSquare:":             "\uf098",
	":Twitter:":                 "\uf099",
	":Facebook:":                "\uf09a",
	":GitHub:":                  "\uf09b",
	":Unlock:":                  "\uf09c",
	":CreditCard:":              "\uf09d",
	":RSS:":                     "\uf09e",
	":HDD:":                     "\uf0a0",
	":Bullhorn:":                "\uf0a1",
	":Bell:":                    "\uf0a2",
	":Certificate:":             "\uf0a3",
	":HandOutlinedRight:":       "\uf0a4",
	":HandOutlinedLeft:":        "\uf0a5",
	":HandOutlinedUp:":          "\uf0a6",
	":HandOutlinedDown:":        "\uf0a7",
	":ArrowCircleLeft:":         "\uf0a8",
	":ArrowCircleRight:":        "\uf0a9",
	":ArrowCircleUp:":           "\uf0aa",
	":ArrowCircleDown:":         "\uf0ab",
	":Globe:":                   "\uf0ac",
	":Wrench:":                  "\uf0ad",
	":Tasks:":                   "\uf0ae",
	":Filter:":                  "\uf0b0",
	":Briefcase:":               "\uf0b1",
	":ArrowsAlt:":               "\uf0b2",
	":Users:":              	 "\uf0c0",
	":Link:":            	     "\uf0c1",
	":Cloud:":          		 "\uf0c2",
	":Flask:":        	         "\uf0c3",
	":Scissors:":      			 "\uf0c4",
	":FilesOutlined:":    		 "\uf0c5",
	":Paperclip:":  			 "\uf0c6",
	":FloppyOutlined:": 		 "\uf0c7",
	":Square:": 				 "\uf0c8",
	":Bars:":                  	 "\uf0c9",
	":ListUl:":                  "\uf0ca",
	":ListOl:":                  "\uf0cb",
	":Strikethrough:":           "\uf0cc",
	":Underline:":          	 "\uf0cd",
	":Table:":        			 "\uf0ce",
	":Magic:":      			 "\uf0d0",
	":Truck:":    				 "\uf0d1",
	":Pinterest:":  			 "\uf0d2",
	":PinterestSquare:":         "\uf0d3",
	":GooglePlusSquare:":        "\uf0d4",
	":GooglePlus:":       		 "\uf0d5",
	":Money:":     				 "\uf0d6",
	":CaretDown:":   			 "\ud0d7",
	":CaretUp:": 				 "\uf0d8",
	":CaretLeft:":               "\uf0d9",
	":CaretRight:":              "\uf0da",
	":Columns:":              	 "\uf0db",
	":Sort:":           		 "\uf0dc",
	":SortDescending:":          "\uf0dd",
	":SortAscending:":       	 "\uf0de",
	":Envelope:":     			 "\uf0e0",
	":LinkedIn:":   			 "\uf0e1",
	":Undo:": 					 "\uf0e2",
	":Gavel:":                   "\uf0e3",
	":Tachometer:":              "\uf0e4",
	":CommentO:":                "\uf0e5",
	":CommentsO:":            	 "\uf0e6",
	":LightningBolt:":           "\uf0e7",
	":Sitemap:":        		 "\uf0e8",
	":Umbrella:":      			 "\uf0e9",
	":Clipboard:":    			 "\uf0ea",
	":LightbulbOutlined:":  	 "\uf0eb",
	":Exchange:": 				 "\uf0ec",
	":CloudDownload:": 			 "\uf0ed",
	":CloudUpload:":             "\uf0ee",
	":UserMd:":             	 "\uf0f0",
	":Stethoscope:":           	 "\uf0f1",
	":Suitcase:":         		 "\uf0f2",
	":BellOutlined:":       	 "\uf0a2",
	":Coffee:":     			 "\uf0f4",
	":Cutlery:":   				 "\uf0f5",
	":FileTextOutlined:": 		 "\uf0f6",
	":BuildingOutlined:":        "\uf0f7",
	":HospitalOutlined:":        "\uf0f8",
	":Ambulance:":      		 "\uf0f9",
	":Medkit:":    			     "\uf0fa",
	":FighterJet:":  			 "\uf0fb",
	":Beer:": 					 "\uf0fc",
	":HSquare:":                 "\uf0fd",
	":PlusSquare:":              "\uf0fe",
	":AngleDoubleLeft:":         "\uf100",
	":AngleDoubleRight:":        "\uf101",
	":AngleDoubleUp:":    	     "\uf102",
	":AngleDoubleDown:":    	 "\uf103",
	":AngleLeft:":      		 "\uf104",
	":AngleRight:": 			 "\uf105",
	":AngleUp:":                 "\uf106",
	":AngleDown:":               "\uf107",
	":Desktop:":                 "\uf108",
	":Laptop:":                  "\uf109",
	":Tablet:":           	     "\uf10a",
	":MobilePhone:":         	 "\uf10b",
	":CircleOutlined:":       	 "\uf10c",
	":QuoteLeft:":     			 "\uf10d",
	":QuoteRight:":   			 "\uf10e",
	":Spinner:": 				 "\uf110",
	":Circle:":                  "\uf111",
	":Reply:":                   "\uf112",
	":GitHubAlt:":               "\uf113",
	":FolderOutlined:":          "\uf114",
	":FolderOpenOutlined:":      "\uf115",
	":SmileOutlined:": "\uf118",
	":FrownOutlined:": "\uf119",
	":MehOutlined:": "\uf11a",
	":Gamepad:": "\uf11b",
	":KeyboardOutlined:": "\uf11c",
	":FlagOutlined:": "\uf11d",
	":FlagCheckered:": "\uf11e",
	":Terminal:": "\uf120",
	":Code:": "\uf121",
	":ReplyAll:": "\uf122",
	":StarHalfOutlined:": "\uf123",
	":LocationArrow:": "\uf124",
	":Crop:": "\uf125",
	":CodeFork:": "\uf126",
	":ChainBroken:": "\uf127",
	":Question:": "\uf128",
	":Info:": "\uf129",
	":Exclamation:": "\uf12a",
	":SuperScript:": "\uf12b",
	":SubScript:": "\uf12c",
	":Eraser:": "\uf12d",
	":PuzzlePiece:": "\uf12e",
	":Microphone:": "\uf130",
	":MicrophoneSlash:": "\uf131",
	":Shield:": "\uf132",
	":CalenderO:": "\uf133",
	":FireExtinguisher:": "\uf134",
	":Rocket:": "\uf135",
}

/*
O-------------------------------O
O License: MIT | Copyright 2020 O
O-------------------------------O
*/
