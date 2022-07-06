package aboutArray

//解法一 排序
// Arrays.sort(nums);
// for(int i = 0; i < nums.length; i++){
//     if(nums[i] != i) return i;
// }
// return nums.length;

//解法二 哈希表
// Set<Integer> set = new HashSet<>();
// for(int i = 0; i < nums.length; i++) set.add(nums[i]);
// for(int i = 0; i <= nums.length; i++)
//     if(!set.contains(i)) return i;
// return -1;

//解法三 位运算
// int res = nums.length;
// for(int i = 0; i < nums.length; i++){
//     res ^= nums[i] ^ i;
// }
// return res;

func missingNumber(nums []int) int {
	length := len(nums)
	for i := 0; i < len(nums); i++ {
		length ^= nums[i]
		length ^= i
	}
	return length
}
