package main

import (
	"fmt"
	"math/rand"
	"time"
)

type module []string

var logicalThinking = module{"https://takeuforward.org/strivers-a2z-dsa-course/must-do-pattern-problems-before-starting-dsa/"}

var basicMaths = module{
	"https://takeuforward.org/data-structure/count-digits-in-a-number/",
	"https://takeuforward.org/maths/reverse-digits-of-a-number",
	"https://takeuforward.org/data-structure/check-if-a-number-is-palindrome-or-not/",
	"https://takeuforward.org/data-structure/find-gcd-of-two-numbers/",
	"https://takeuforward.org/maths/check-if-a-number-is-armstrong-number-or-not/",
	"https://takeuforward.org/data-structure/print-all-divisors-of-a-given-number/",
	"https://takeuforward.org/data-structure/check-if-a-number-is-prime-or-not/",
}

var recursion = module{
	"https://takeuforward.org/recursion/print-name-n-times-using-recursion/",
	"https://takeuforward.org/recursion/print-1-to-n-using-recursion/",
	"https://takeuforward.org/recursion/print-n-to-1-using-recursion/",
	"https://takeuforward.org/data-structure/sum-of-first-n-natural-numbers/",
	"https://takeuforward.org/data-structure/factorial-of-a-number-iterative-and-recursive/",
	"https://takeuforward.org/data-structure/reverse-a-given-array/",
	"https://takeuforward.org/data-structure/check-if-the-given-string-is-palindrome-or-not/",
	"https://takeuforward.org/arrays/print-fibonacci-series-up-to-nth-term/",
}

var hashing = module{
	"https://takeuforward.org/data-structure/count-frequency-of-each-element-in-the-array/",
	"https://takeuforward.org/arrays/find-the-highest-lowest-frequency-element/",
}

var sorting = module{
	"https://takeuforward.org/sorting/selection-sort-algorithm/",
	"https://takeuforward.org/data-structure/bubble-sort-algorithm/",
	"https://takeuforward.org/data-structure/insertion-sort-algorithm/",
	"https://takeuforward.org/data-structure/merge-sort-algorithm/",
	"https://takeuforward.org/arrays/recursive-bubble-sort-algorithm/",
	"https://takeuforward.org/arrays/recursive-insertion-sort-algorithm/",
	"https://takeuforward.org/data-structure/quick-sort-algorithm/",
}

var easyArrays = module{
	"https://takeuforward.org/data-structure/find-the-largest-element-in-an-array/",
	"https://takeuforward.org/data-structure/find-second-smallest-and-second-largest-element-in-an-array/",
	"https://takeuforward.org/data-structure/check-if-an-array-is-sorted/",
	"https://takeuforward.org/data-structure/remove-duplicates-in-place-from-sorted-array/",
	"https://takeuforward.org/data-structure/left-rotate-the-array-by-one/",
	"https://takeuforward.org/data-structure/rotate-array-by-k-elements/",
	"https://takeuforward.org/data-structure/move-all-zeros-to-the-end-of-the-array/",
	"https://takeuforward.org/data-structure/linear-search-in-c/",
	"https://takeuforward.org/data-structure/union-of-two-sorted-arrays/",
	"https://takeuforward.org/arrays/find-the-missing-number-in-an-array/",
	"https://takeuforward.org/data-structure/count-maximum-consecutive-ones-in-the-array/",
	"https://takeuforward.org/arrays/find-the-number-that-appears-once-and-the-other-numbers-twice/",
	"https://takeuforward.org/data-structure/longest-subarray-with-given-sum-k/",
	"https://takeuforward.org/arrays/longest-subarray-with-sum-k-postives-and-negatives/",
}

var mediumArrays = module{
	"https://takeuforward.org/data-structure/two-sum-check-if-a-pair-with-given-sum-exists-in-array/",
	"https://takeuforward.org/data-structure/sort-an-array-of-0s-1s-and-2s/",
	"https://takeuforward.org/data-structure/find-the-majority-element-that-occurs-more-than-n-2-times/",
	"https://takeuforward.org/data-structure/kadanes-algorithm-maximum-subarray-sum-in-an-array/",
	"https://takeuforward.org/data-structure/kadanes-algorithm-maximum-subarray-sum-in-an-array/",
	"https://takeuforward.org/data-structure/stock-buy-and-sell/",
	"https://takeuforward.org/arrays/rearrange-array-elements-by-sign/",
	"https://takeuforward.org/data-structure/next_permutation-find-next-lexicographically-greater-permutation/",
	"https://takeuforward.org/data-structure/leaders-in-an-array/",
	"https://takeuforward.org/data-structure/longest-consecutive-sequence-in-an-array/",
	"https://takeuforward.org/data-structure/set-matrix-zero/",
	"https://takeuforward.org/data-structure/rotate-image-by-90-degree/",
	"https://takeuforward.org/data-structure/spiral-traversal-of-matrix/",
	"https://takeuforward.org/arrays/count-subarray-sum-equals-k/",
	"https://takeuforward.org/plus/dsa/problems/prime-factorisation-of-a-number",
}

var hardArrays = module{
	"https://takeuforward.org/data-structure/program-to-generate-pascals-triangle/",
	"https://takeuforward.org/data-structure/majority-elementsn-3-times-find-the-elements-that-appears-more-than-n-3-times-in-the-array/",
	"https://takeuforward.org/data-structure/3-sum-find-triplets-that-add-up-to-a-zero/",
	"https://takeuforward.org/data-structure/4-sum-find-quads-that-add-up-to-a-target-value/",
	"https://takeuforward.org/data-structure/length-of-the-longest-subarray-with-zero-sum/",
	"https://takeuforward.org/data-structure/count-the-number-of-subarrays-with-given-xor-k/",
	"https://takeuforward.org/data-structure/merge-overlapping-sub-intervals/",
	"https://takeuforward.org/data-structure/merge-two-sorted-arrays-without-extra-space/",
	"https://takeuforward.org/data-structure/find-the-repeating-and-missing-numbers/",
	"https://takeuforward.org/data-structure/count-inversions-in-an-array/",
	"https://takeuforward.org/data-structure/count-reverse-pairs/",
	"https://takeuforward.org/data-structure/maximum-product-subarray-in-an-array/",
}

var bs1d = module{
	"https://takeuforward.org/data-structure/binary-search-explained/",
	"https://takeuforward.org/arrays/implement-lower-bound-bs-2/",
	"https://takeuforward.org/arrays/implement-upper-bound/",
	"https://takeuforward.org/arrays/search-insert-position/",
	"https://takeuforward.org/arrays/floor-and-ceil-in-sorted-array/",
	"https://takeuforward.org/data-structure/last-occurrence-in-a-sorted-array/",
	"https://takeuforward.org/data-structure/count-occurrences-in-sorted-array/",
	"https://takeuforward.org/data-structure/search-element-in-a-rotated-sorted-array/",
	"https://takeuforward.org/arrays/search-element-in-rotated-sorted-array-ii/",
	"https://takeuforward.org/data-structure/minimum-in-rotated-sorted-array/",
	"https://takeuforward.org/arrays/find-out-how-many-times-the-array-has-been-rotated/",
	"https://takeuforward.org/data-structure/search-single-element-in-a-sorted-array/",
	"https://takeuforward.org/data-structure/peak-element-in-array/",
}

var bsOnAnswers = module{
	"https://takeuforward.org/binary-search/finding-sqrt-of-a-number-using-binary-search/",
	"https://takeuforward.org/data-structure/nth-root-of-a-number-using-binary-search/",
	"https://takeuforward.org/binary-search/koko-eating-bananas/",
	"https://takeuforward.org/arrays/minimum-days-to-make-m-bouquets/",
	"https://takeuforward.org/arrays/find-the-smallest-divisor-given-a-threshold/",
	"https://takeuforward.org/arrays/capacity-to-ship-packages-within-d-days/",
	"https://takeuforward.org/arrays/kth-missing-positive-number/",
	"https://takeuforward.org/data-structure/aggressive-cows-detailed-solution/",
	"https://takeuforward.org/data-structure/allocate-minimum-number-of-pages/",
	"https://takeuforward.org/arrays/split-array-largest-sum/",
	"https://takeuforward.org/arrays/painters-partition-problem/",
}

var bsOn2d = module{
	"https://takeuforward.org/arrays/find-the-row-with-maximum-number-of-1s/",
	"https://takeuforward.org/data-structure/search-in-a-sorted-2d-matrix/",
	"https://takeuforward.org/arrays/search-in-a-row-and-column-wise-sorted-matrix/",
	"https://takeuforward.org/binary-search/find-peak-element-ii",
	"https://takeuforward.org/data-structure/median-of-row-wise-sorted-matrix/",
}

var basicStrings = module{
	"https://leetcode.com/problems/remove-outermost-parentheses/description/",
	"https://takeuforward.org/data-structure/reverse-words-in-a-string/",
	"https://leetcode.com/problems/largest-odd-number-in-string/description/",
	"https://leetcode.com/problems/longest-common-prefix/description/?source=submission-ac",
	"https://leetcode.com/problems/isomorphic-strings/description/",
}

var mediumStrings = module{
	"https://leetcode.com/problems/rotate-string/",
	"https://takeuforward.org/data-structure/check-if-two-strings-are-anagrams-of-each-other/",
	"https://leetcode.com/problems/sort-characters-by-frequency/description/",
	"https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/description/",
	"https://leetcode.com/problems/roman-to-integer/description/",
	"https://leetcode.com/problems/string-to-integer-atoi/description/",
}

var ll = module{
	"https://takeuforward.org/linked-list/insert-at-the-head-of-a-linked-list",
	"https://takeuforward.org/data-structure/delete-last-node-of-linked-list/",
	"https://takeuforward.org/linked-list/find-the-length-of-a-linked-list",
	"https://takeuforward.org/linked-list/search-an-element-in-a-linked-list",
	"https://takeuforward.org/data-structure/find-middle-element-in-a-linked-list/",
	"https://takeuforward.org/data-structure/reverse-a-linked-list/",
	"https://takeuforward.org/data-structure/detect-a-cycle-in-a-linked-list/",
	"https://takeuforward.org/data-structure/starting-point-of-loop-in-a-linked-list/",
	"https://takeuforward.org/linked-list/length-of-loop-in-linked-list",
	"https://takeuforward.org/data-structure/check-if-given-linked-list-is-plaindrome/",
	"https://takeuforward.org/data-structure/segregate-even-and-odd-nodes-in-linkedlist",
	"https://takeuforward.org/data-structure/remove-n-th-node-from-the-end-of-a-linked-list/",
	"https://takeuforward.org/linked-list/delete-the-middle-node-of-the-linked-list",
	"https://takeuforward.org/linked-list/sort-a-linked-list",
	"https://takeuforward.org/plus/dsa/problems/sort-a-ll-of-0's-1's-and-2's",
	"https://takeuforward.org/data-structure/find-intersection-of-two-linked-lists/",
	"https://www.geeksforgeeks.org/problems/add-1-to-a-number-represented-as-linked-list/1",
	"https://takeuforward.org/data-structure/add-two-numbers-represented-as-linked-lists/",
}

var dll = module{
	"https://takeuforward.org/data-structure/insert-at-end-of-doubly-linked-list/",
	"https://takeuforward.org/data-structure/delete-last-node-of-a-doubly-linked-list/",
	"https://takeuforward.org/data-structure/reverse-a-doubly-linked-list/",
	"https://takeuforward.org/plus/dsa/problems/delete-all-occurrences-of-a-key-in-dll",
	"https://takeuforward.org/plus/dsa/problems/find-pairs-with-given-sum-in-doubly-linked-list",
	"https://takeuforward.org/plus/dsa/problems/remove-duplicated-from-sorted-dll",
}

var recursionPatternWise = module{
	"https://takeuforward.org/data-structure/implement-powxn-x-raised-to-the-power-n/",
	"https://leetcode.com/problems/count-good-numbers/description/",
	"https://takeuforward.org/plus/dsa/problems/generate-binary-strings-without-consecutive-1s",
	"https://leetcode.com/problems/generate-parentheses/description/",
	"https://takeuforward.org/data-structure/power-set-print-all-the-possible-subsequences-of-the-string/",
	"https://takeuforward.org/plus/dsa/problems/count-all-subsequences-with-sum-k",
	"https://takeuforward.org/plus/dsa/problems/check-if-there-exists-a-subsequence-with-sum-k",
	"https://takeuforward.org/data-structure/combination-sum-1/",
	"https://takeuforward.org/data-structure/combination-sum-ii-find-all-unique-combinations/",
	"https://takeuforward.org/data-structure/subset-sum-sum-of-all-subsets/",
	"https://takeuforward.org/data-structure/subset-ii-print-all-the-unique-subsets/",
	"https://leetcode.com/problems/combination-sum-iii/description/",
	"https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/",
}

var recursionHard = module{
	"https://takeuforward.org/data-structure/palindrome-partitioning/",
	"https://takeuforward.org/data-structure/word-search-leetcode/",
	"https://takeuforward.org/data-structure/n-queen-problem-return-all-distinct-solutions-to-the-n-queens-puzzle/",
	"https://takeuforward.org/data-structure/rat-in-a-maze/",
	"https://takeuforward.org/data-structure/sudoku-solver/",
}

var bitManipulation = module{
	"https://takeuforward.org/plus/dsa/problems/check-if-the-i-th-bit-is-set-or-not",
	"https://takeuforward.org/plus/dsa/problems/check-if-a-number-is-odd-or-not",
	"https://leetcode.com/problems/power-of-two/description/",
	"https://takeuforward.org/plus/dsa/problems/count-the-number-of-set-bits",
	"https://takeuforward.org/plus/dsa/problems/swap-two-numbers",
	"https://www.geeksforgeeks.org/dsa/set-rightmost-unset-bit/",
}

var mediumBitManipulation = module{
	"https://leetcode.com/problems/minimum-bit-flips-to-convert-number/description/",
	"https://leetcode.com/problems/single-number/description/",
	"https://leetcode.com/problems/subsets/description/",
	"https://takeuforward.org/plus/dsa/problems/xor-of-numbers-in-a-given-range",
	"https://takeuforward.org/plus/dsa/problems/single-number---iii",
	"https://www.youtube.com/watch?v=LT7XhVdeRyg",
	"https://leetcode.com/problems/count-primes/description/",
	"https://leetcode.com/problems/powx-n/description/",
}

var stackLearning = module{
	// "https://takeuforward.org/data-structure/implement-stack-using-array/",
	// "https://takeuforward.org/data-structure/implement-queue-using-array/",
	// "https://takeuforward.org/data-structure/implement-stack-using-single-queue/",
	// "https://takeuforward.org/data-structure/implement-queue-using-stack/",
	// "https://takeuforward.org/data-structure/implement-stack-using-linked-list/",
	// "https://takeuforward.org/data-structure/implement-queue-using-linked-list/",
	"https://takeuforward.org/data-structure/check-for-balanced-parentheses/",
	"https://takeuforward.org/data-structure/implement-min-stack-o2n-and-on-space-complexity/",
	"https://takeuforward.org/data-structure/infix-to-postfix/",
	"https://takeuforward.org/data-structure/infix-to-prefix/",
	"postfix-to-infix: ab-de+f*/   ",
	"https://takeuforward.org/plus/dsa/problems/prefix-to-infix-conversion",
	"https://takeuforward.org/plus/dsa/problems/postfix-to-prefix-conversion",
	"https://takeuforward.org/plus/dsa/problems/prefix-to-postfix-conversion",
}

var stackPart2 = module{
	"https://takeuforward.org/data-structure/next-greater-element-using-stack/",
	"https://takeuforward.org/plus/dsa/problems/next-smaller-element",
	"https://takeuforward.org/data-structure/trapping-rainwater/",
	"https://leetcode.com/problems/sum-of-subarray-minimums/description/",
	"https://leetcode.com/problems/asteroid-collision/description/",
	"https://leetcode.com/problems/sum-of-subarray-ranges/description/",
	"https://leetcode.com/problems/remove-k-digits/",
	"https://takeuforward.org/data-structure/area-of-largest-rectangle-in-histogram/",
	"https://leetcode.com/problems/maximal-rectangle/description/",
	"https://takeuforward.org/data-structure/sliding-window-maximum/",
	"https://leetcode.com/problems/online-stock-span/description/",
	"https://takeuforward.org/plus/dsa/problems/celebrity-problem",
	"https://leetcode.com/problems/lru-cache/description/",
}

var slidingWindow = module{
	"https://takeuforward.org/data-structure/length-of-longest-substring-without-any-repeating-character/",
	"https://leetcode.com/problems/max-consecutive-ones-iii/description/",
	"https://takeuforward.org/plus/dsa/problems/fruit-into-baskets",
	"https://leetcode.com/problems/longest-repeating-character-replacement/description/",
	"https://leetcode.com/problems/binary-subarrays-with-sum/description/",
	"https://leetcode.com/problems/count-number-of-nice-subarrays/description/",
}

var Heaps = module{
	"https://takeuforward.org/data-structure/introduction-to-priority-queues-using-binary-heaps",
	"https://takeuforward.org/plus/dsa/problems/check-if-an-array-represents-a-min-heap-",
	"https://takeuforward.org/plus/dsa/problems/convert-min-heap-to-max-heap",
	"https://takeuforward.org/data-structure/kth-largest-smallest-element-in-an-array/",
	"https://www.geeksforgeeks.org/dsa/nearly-sorted-algorithm/",
	"https://takeuforward.org/data-structure/replace-elements-by-its-rank-in-the-array/",
}

var greedy = module{
	"https://takeuforward.org/data-structure/assign-cookies",
	"https://takeuforward.org/Greedy/lemonade-change",
}

var binaryTrees = module{
	"https://leetcode.com/problems/binary-tree-preorder-traversal/description/",
	"https://leetcode.com/problems/binary-tree-inorder-traversal/description/",
	"https://leetcode.com/problems/binary-tree-postorder-traversal/description/",
	"https://leetcode.com/problems/binary-tree-level-order-traversal/description/",
	"https://leetcode.com/problems/binary-tree-preorder-traversal/description/",
	"https://leetcode.com/problems/binary-tree-inorder-traversal/description/",
	"https://leetcode.com/problems/binary-tree-postorder-traversal/description/",
}

var modules = []module{
	// basicMaths,
	// recursion,
	// hashing,
	// easyArrays,
	bs1d,
	mediumArrays,
	hardArrays,
	basicStrings,
	bsOn2d,
	bitManipulation,
	sorting,
	ll,
	dll,
}

var newOnes = []module{
	bsOnAnswers,
	mediumStrings,
	recursionPatternWise,
	recursionHard,
	mediumBitManipulation,
	stackLearning,
	stackPart2,
	slidingWindow,
}

func main() {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	// fmt.Println("[Pattern]", logicalThinking[0], r.Intn(22)+1)

	selected := r.Perm(len(modules))[:3]
	for _, i := range selected {
		fmt.Println("[OLD]", modules[i][r.Intn(len(modules[i]))])
	}

	for i := range newOnes {
		fmt.Println("[NEW]", newOnes[i][r.Intn(len(newOnes[i]))])
	}
}

// to do
// KMP algo - to search for substring
// Rabin-Karp String Matching Algorithm
