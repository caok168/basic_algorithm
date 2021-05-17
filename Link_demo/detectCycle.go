// 环路检测

func detectCycle(head *ListNode) *ListNode {
    seen := map[*ListNode]struct{}{}
    for head != nil {
        if _, ok := seen[head]; ok {
            return head
        }
        seen[head] = struct{}{}
        head = head.Next
    }
    return nil
}

