package lcs

// func checkIfScramble(s1 string, s2 string, m map[string]bool) bool {
// 	if s1 == s2 {
// 		return true
// 	}

// 	if len(s1) <= 1 || len(s2) <= 1 {
// 		return false
// 	}

// 	if _, ok := m[s1+":"+s2]; ok {
// 		return m[s1+":"+s2]
// 	}

// 	var b1, b2, b3, b4 bool
// 	var p1, p2, p3, p4 string
// 	for i := 0; i < len(s1)-1; i++ {
// 		p1 = s1[0:i+1] + ":" + s2[len(s2)-(i+1):]
// 		p2 = s1[i+1:] + ":" + s2[0:len(s2)-(i+1)]
// 		p3 = s1[0:i+1] + ":" + s2[0:i+1]
// 		p4 = s1[i+1:] + ":" + s2[i+1:]

// 		if _, ok := m[p1]; ok {
// 			b1 = m[p1]
// 		} else {
// 			b1 = checkIfScramble(s1[0:i+1], s2[len(s2)-(i+1):], m)
// 		}
// 		if b1 {
// 			if _, ok := m[p2]; ok {
// 				b2 = m[p2]
// 			} else {
// 				b2 = checkIfScramble(s1[i+1:], s2[0:len(s2)-(i+1)], m)
// 			}
// 		}
// 		if _, ok := m[p3]; ok {
// 			b3 = m[p3]
// 		} else {
// 			b3 = checkIfScramble(s1[0:i+1], s2[0:i+1], m)
// 		}

// 		if b3 {
// 			if _, ok := m[p4]; ok {
// 				b4 = m[p4]
// 			} else {
// 				b4 = checkIfScramble(s1[i+1:], s2[i+1:], m)
// 			}
// 		}

// 		if (b1 && b2) || (b3 && b4) {
// 			m[s1+":"+s2] = true
// 			return true
// 		}
// 	}

// 	m[s1+":"+s2] = false
// 	return false
// }
