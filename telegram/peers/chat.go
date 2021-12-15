package peers

import (
	"context"

	"github.com/go-faster/errors"

	"github.com/gotd/td/constant"
	"github.com/gotd/td/tg"
)

// Chat is chat peer.
type Chat struct {
	raw *tg.Chat
	m   *Manager
}

// Chat creates new Chat, attached to this manager.
func (m *Manager) Chat(u *tg.Chat) Chat {
	return Chat{
		raw: u,
		m:   m,
	}
}

// GetChat gets Chat using given id.
func (m *Manager) GetChat(ctx context.Context, id int64) (Chat, error) {
	ch, err := m.getChat(ctx, id)
	if err != nil {
		return Chat{}, err
	}
	return m.Chat(ch), nil
}

// Raw returns raw *tg.Chat.
func (c Chat) Raw() *tg.Chat {
	return c.raw
}

// ID returns entity ID.
func (c Chat) ID() int64 {
	return c.raw.GetID()
}

// TDLibPeerID returns TDLibPeerID for this entity.
func (c Chat) TDLibPeerID() (r constant.TDLibPeerID) {
	r.Chat(c.raw.GetID())
	return r
}

// VisibleName returns visible name of peer.
//
// It returns FirstName + " " + LastName for users, and title for chats and channels.
func (c Chat) VisibleName() string {
	return c.raw.GetTitle()
}

// Username returns peer username, if any.
func (c Chat) Username() (string, bool) {
	return "", false
}

// Restricted whether this user/chat/channel is restricted.
func (c Chat) Restricted() ([]tg.RestrictionReason, bool) {
	return nil, false
}

// Verified whether this user/chat/channel is verified by Telegram.
func (c Chat) Verified() bool {
	return false
}

// Scam whether this user/chat/channel is probably a scam.
func (c Chat) Scam() bool {
	return false
}

// Fake whether this user/chat/channel was reported by many users as a fake or scam: be
// careful when interacting with it.
func (c Chat) Fake() bool {
	return false
}

// InputPeer returns input peer for this peer.
func (c Chat) InputPeer() tg.InputPeerClass {
	return c.raw.AsInputPeer()
}

// Sync updates current object.
func (c Chat) Sync(ctx context.Context) error {
	raw, err := c.m.updateChat(ctx, c.raw.ID)
	if err != nil {
		return errors.Wrap(err, "get chat")
	}
	*c.raw = *raw
	return nil
}

// Report reports a peer for violation of telegram's Terms of Service.
func (c Chat) Report(ctx context.Context, reason tg.ReportReasonClass, message string) error {
	if _, err := c.m.api.AccountReportPeer(ctx, &tg.AccountReportPeerRequest{
		Peer:    c.InputPeer(),
		Reason:  reason,
		Message: message,
	}); err != nil {
		return errors.Wrap(err, "report")
	}
	return nil
}

// Photo returns peer photo, if any.
func (c Chat) Photo(ctx context.Context) (*tg.Photo, bool, error) {
	full, err := c.m.getChatFull(ctx, c.raw.GetID())
	if err != nil {
		return nil, false, errors.Wrap(err, "get full chat")
	}

	chatPhoto, ok := full.GetChatPhoto()
	if !ok {
		return nil, false, nil
	}

	p, ok := chatPhoto.AsNotEmpty()
	return p, ok, nil
}

// Creator whether the current user is the creator of this group.
func (c Chat) Creator() bool {
	return c.raw.GetCreator()
}

// Kicked the current user was kicked from the group.
func (c Chat) Kicked() bool {
	return c.raw.GetKicked()
}

// Left whether the current user has left this group.
func (c Chat) Left() bool {
	return c.raw.GetLeft()
}

// Deactivated whether the group was migrated.
func (c Chat) Deactivated() bool {
	return c.raw.GetDeactivated()
}

// CallActive whether a group call or livestream is currently active.
func (c Chat) CallActive() bool {
	return c.raw.GetCallActive()
}

// CallNotEmpty whether there's anyone in the group call or livestream.
func (c Chat) CallNotEmpty() bool {
	return c.raw.GetCallNotEmpty()
}

// NoForwards whether that message forwarding from this channel is not allowed.
func (c Chat) NoForwards() bool {
	return c.raw.GetNoforwards()
}

// MigratedTo returns a supergroup to which this chat migrated.
func (c Chat) MigratedTo() (tg.InputChannelClass, bool) {
	return c.raw.GetMigratedTo()
}

// ParticipantsCount returns count of participants.
func (c Chat) ParticipantsCount() int {
	return c.raw.GetParticipantsCount()
}

// AdminRights returns admin rights of the user in this channel.
//
// See https://core.telegram.org/api/rights.
func (c Chat) AdminRights() (tg.ChatAdminRights, bool) {
	return c.raw.GetAdminRights()
}

// DefaultBannedRights returns default chat rights.
//
// See https://core.telegram.org/api/rights.
func (c Chat) DefaultBannedRights() (tg.ChatBannedRights, bool) {
	return c.raw.GetDefaultBannedRights()
}

// Leave leaves this chat.
//
// Parameter deleteMyMessages denotes to remove the entire chat history of this user in this chat.
func (c Chat) Leave(ctx context.Context, deleteMyMessages bool) error {
	_, err := c.m.api.MessagesDeleteChatUser(ctx, &tg.MessagesDeleteChatUserRequest{
		RevokeHistory: deleteMyMessages,
		ChatID:        c.raw.GetID(),
		UserID:        &tg.InputUserSelf{},
	})
	return err
}

// TODO(tdakkota): add more getters, helpers and convertors
