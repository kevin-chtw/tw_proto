package cproto

import "testing"

func TestGameSeatMessages(t *testing.T) {
	req := &GameSeatReq{Seat: -1}
	if req.GetSeat() != -1 {
		t.Fatalf("seat=%d", req.GetSeat())
	}
	ack := &GameSeatAck{Uid: "u", Seat: 2, Ready: false}
	if ack.GetUid() != "u" || ack.GetSeat() != 2 || ack.GetReady() {
		t.Fatalf("ack=%+v", ack)
	}
	enter := &EnterGameAck{SeatMode: 1}
	if enter.GetSeatMode() != 1 {
		t.Fatalf("seat_mode=%d", enter.GetSeatMode())
	}
}
