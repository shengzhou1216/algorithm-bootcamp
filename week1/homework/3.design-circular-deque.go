package week1
type MyCircularDeque struct {
    cap int
    data []int
    front int
    rear int
}


func Constructor(k int) MyCircularDeque {
    return MyCircularDeque{
        cap: k,
        data: make([]int,k),
        front: -1,
        rear: 0,
    }
}


func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.IsFull(){
        return false
    }
    if this.front == -1 {
        this.front = 0
        this.rear = 0
    }else if this.front == 0 {
        this.front = this.cap - 1
    }else{
        this.front--
    }
    this.data[this.front] = value
    return true
}


func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.IsFull(){
        return false
    }
    if this.front == -1 {
        this.front = 0
        this.rear = 0
    }else if this.rear == this.cap - 1 {
        this.rear = 0
    }else {
        this.rear++
    }
    this.data[this.rear] = value
    return true
}


func (this *MyCircularDeque) DeleteFront() bool {
    if this.IsEmpty(){
        return false
    }
    if this.front == this.rear {
        this.front = -1
        this.rear = -1
    }else if this.front == this.cap - 1 {
        this.front = 0
    }else {
        this.front++
    }
    return true
}


func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty(){
        return false
    }
    if this.front == this.rear {
        this.front = -1
        this.rear = -1
    }else if this.rear == 0 {
        this.rear = this.cap - 1
    }else{
        this.rear--
    }
    return true
}


func (this *MyCircularDeque) GetFront() int {
    if this.IsEmpty(){
        return -1
    }
    return this.data[this.front]
}


func (this *MyCircularDeque) GetRear() int {
    if this.IsEmpty() || this.rear < 0 {
        return -1
    }
    return this.data[this.rear]
}


func (this *MyCircularDeque) IsEmpty() bool {
    return this.front == -1
}


func (this *MyCircularDeque) IsFull() bool {
    return (this.front == this.rear + 1)  || (this.front == 0 && this.rear == this.cap - 1)
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */