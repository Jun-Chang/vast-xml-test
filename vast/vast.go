package vast

import "encoding/xml"

type Vast struct {
	XMLName xml.Name `xml:"VAST"`
	Version string   `xml:"version,attr"`
	Ads     []Ad     `xml:"Ad"`
}

type Ad struct {
	InLine *InLine `xml:",omitempty"`
	ID     int     `xml:"id,attr"`
}

type InLine struct {
	AdSystem    AdSystem    `xml:"AdSystem"`
	AdTitle     string      `xml:"AdTitle"`
	Impression  URIWithID   `xml:"Impression"`
	Description string      `xml:"Description"`
	Error       CData       `xml:"Error"`
	Extensions  []Extension `xml:"Extensions>Extension"`
	Creatives   []Creative  `xml:"Creatives>Creative"`
}

type AdSystem struct {
	Name    string `xml:",chardata"`
	Version string `xml:"version,attr"`
}

type URIWithID struct {
	URI string `xml:",cdata"`
	ID  string `xml:"id,attr"`
}

type CData struct {
	Text string `xml:",cdata"`
}

type Extension struct{}

type Creative struct {
	Linear *Linear `xml:"Linear,omitempty"`
}

type Linear struct {
	Duration       string      `xml:"Duration"`
	MediaFiles     []MediaFile `xml:"MediaFiles>MediaFile"`
	VideoClicks    VideoClicks `xml:"VideoClicks"`
	TrackingEvents []Tracking  `xml:"TrackingEvents>Tracking"`
}

type VideoClicks struct {
	ClickThrough  URIWithID `xml:"ClickThrough"`
	ClickTracking URIWithID `xml:"ClickTracking"`
	CustomClick   URIWithID `xml:"CustomClick"`
}

type Tracking struct {
	URI    string `xml:",cdata"`
	Event  string `xml:"event,attr"`
	Offset string `xml:"offset,attr,omitempty"`
}

type MediaFile struct {
	URI      string `xml:",cdata"`
	ID       string `xml:"id,attr"`
	Delivery string `xml:"delivery,attr"`
	Type     string `xml:"type,attr"`
	Width    int    `xml:"width,attr"`
	Height   int    `xml:"height,attr"`
}

func New() *Vast {
	return &Vast{
		Version: "3.0",
		Ads: []Ad{
			{
				InLine: &InLine{
					AdSystem: AdSystem{
						Name:    "vast sample",
						Version: "0.1",
					},
					AdTitle: "title",
					Impression: URIWithID{
						URI: "http://xxx/impression",
						ID:  "1",
					},
					Description: "description",
					Error: CData{
						Text: "http://xxx/error",
					},
					Extensions: []Extension{},
					Creatives: []Creative{
						{
							Linear: &Linear{
								Duration: "00:50:00",
								MediaFiles: []MediaFile{
									{
										URI:      "http://xxx/video.mp4",
										ID:       "1",
										Delivery: "progressive",
										Type:     "video/mp4",
										Width:    640,
										Height:   480,
									},
								},
								VideoClicks: VideoClicks{
									ClickThrough: URIWithID{
										URI: "http://xxx/click_through",
										ID:  "1",
									},
									ClickTracking: URIWithID{
										URI: "http://xxx/click_traking",
										ID:  "1",
									},
									CustomClick: URIWithID{
										URI: "http://xxx/custom_click",
										ID:  "1",
									},
								},
								TrackingEvents: []Tracking{
									{
										URI:   "http://xxx/complete",
										Event: "complete",
									},
								},
							},
						},
					},
				},
				ID: 1,
			},
		},
	}
}
