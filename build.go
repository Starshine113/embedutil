package embedutil

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

// Build returns the actual message embed object
func (e *Embed) Build() *discordgo.MessageEmbed {
	embed, _ := e.BuildReportTrimming()
	return embed
}

// BuildReportTrimming is the same as Build() but returns an additional `true` if one or more fields were trimmed
func (e *Embed) BuildReportTrimming() (embed *discordgo.MessageEmbed, trimmed bool) {
	if len(e.title) > TitleLimit {
		e.title = e.title[:TitleLimit-1]
		trimmed = true
	}
	if len(e.description) > DescriptionLimit {
		e.description = e.description[:DescriptionLimit-1]
		trimmed = true
	}
	if e.author != nil {
		if len(e.author.name) > TitleLimit {
			e.author.name = e.author.name[:TitleLimit-1]
			trimmed = true
		}
	} else {
		e.author = &Author{}
	}
	if e.footer != nil {
		if len(e.footer.text) > FooterTextLimit {
			e.footer.text = e.footer.text[:FooterTextLimit-1]
			trimmed = true
		}
	} else {
		e.footer = &Footer{}
	}
	if len(e.fields) > FieldLimit {
		e.fields = e.fields[:FieldLimit-1]
	}
	for i := len(e.fields); i > 0; i-- {
		if e.totalLen() < EmbedLimit {
			break
		}
		e.fields = e.fields[:i-1]
	}
	var t string
	if e.timestamp != nil {
		t = e.timestamp.Format(time.RFC3339)
	}

	fields := make([]*discordgo.MessageEmbedField, 0)

	for _, f := range e.fields {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   f.name,
			Value:  f.value,
			Inline: f.inline,
		})
	}

	embed = &discordgo.MessageEmbed{
		Title:       e.title,
		Description: e.description,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    e.author.name,
			URL:     e.author.url,
			IconURL: e.author.avatarURL,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text:    e.footer.text,
			IconURL: e.footer.iconURL,
		},
		Fields: fields,
		Color:  e.color,
		Image: &discordgo.MessageEmbedImage{
			URL: e.imageURL,
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: e.thumbnailURL,
		},
		URL:       e.url,
		Timestamp: t,
	}

	return embed, trimmed
}

func (e *Embed) totalLen() (t int) {
	t += len(e.title)
	t += len(e.description)
	if e.author != nil {
		t += len(e.author.name)
	}
	if e.footer != nil {
		t += len(e.footer.text)
	}
	for _, f := range e.fields {
		t += len(f.name)
		t += len(f.value)
	}
	return
}

// BuildSend returns a MessageSend object
func (e *Embed) BuildSend() *discordgo.MessageSend {
	return &discordgo.MessageSend{
		Content: e.content,
		Embed:   e.Build(),
	}
}

// BuildEdit returns a MessageEdit object
func (e *Embed) BuildEdit(channelID, messageID, content string) *discordgo.MessageEdit {
	return discordgo.NewMessageEdit(channelID, messageID).SetContent(content).SetEmbed(e.Build())
}
