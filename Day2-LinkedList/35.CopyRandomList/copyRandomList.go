/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {

    //方法1:使用哈希表进行复制链表,之后再指定新链表的next、random指针的指向
//     if head == nil {return nil}

//     dic := map[*Node]*Node{}
//     //将旧链表的节点存入哈希表中
//     cur := head
//     for cur !=nil{
//         //这里不能直接写dic[cur]=cur;因为主要考察在链表的复制,不能直接取巧
//         dic[cur] = &Node{Val:cur.Val}//结构体在初始化的时候是其内的字段名是可以不全部一次性初始化的;没有写的字段名默认是是其类型的零值
//         cur = cur.Next
//     }

//     //重新制定复制的链表的Next、Random指针
//     cur = head
//     //fmt.Println(dic[cur.Next.Next.Random].Val)
//     for cur !=nil {
//         dic[cur].Next = dic[cur.Next]
//         dic[cur].Random = dic[cur.Random]

//         cur =cur.Next
//     }
//    // fmt.Println(dic)
//   // fmt.Printf("%#v",dic)

//     return dic[head]

       //方法2:拼接+拆分
       //对原链表进行复制并插入原链表中,之后指定新插入链表的Next和random指针,最后将其拆分出两个链表即可;
       if head ==nil{return nil}

       //进行复制链表并将其插入旧链表中;
       cur := head
       //每次只是将当前的节点进行复制;并指定复制后节点的Next指针的指向
       for cur != nil{
           tmp := &Node{Val:cur.Val}
           tmp.Next = cur.Next
           cur.Next = tmp
           cur = tmp.Next
       }

       //指定下插入链表的random指针指向
       cur = head
       for cur != nil{
           if cur.Random != nil{cur.Next.Random = cur.Random.Next}
           cur = cur.Next.Next //注意进行下一节点的指定的时候下一节点的位置是Next.Next 对应原来的Next
       }

       //进行拆分链表
       cur = head.Next
       ret := cur

       pre := head
       for cur.Next != nil{
           pre.Next = pre.Next.Next
           cur.Next = cur.Next.Next
           pre = pre.Next
           cur = cur.Next

       }
       pre.Next = nil
       return ret

}

