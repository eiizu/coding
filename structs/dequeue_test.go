package structs

import "testing"

func TestDeque_PushFront(t *testing.T) {
	type fields struct {
		head *node
		tail *node
	}
	type args struct {
		data int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dq := &Deque{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			dq.PushFront(tt.args.data)
		})
	}
}

func TestDeque_PushBack(t *testing.T) {
	type fields struct {
		head *node
		tail *node
	}
	type args struct {
		data int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dq := &Deque{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			dq.PushBack(tt.args.data)
		})
	}
}

func TestDeque_PopFront(t *testing.T) {
	type fields struct {
		head *node
		tail *node
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dq := &Deque{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			got, err := dq.PopFront()
			if (err != nil) != tt.wantErr {
				t.Errorf("Deque.PopFront() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Deque.PopFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeque_PopBack(t *testing.T) {
	type fields struct {
		head *node
		tail *node
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dq := &Deque{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			got, err := dq.PopBack()
			if (err != nil) != tt.wantErr {
				t.Errorf("Deque.PopBack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Deque.PopBack() = %v, want %v", got, tt.want)
			}
		})
	}
}
