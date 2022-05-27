package leetcodepractice

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	// We need to get the mid point of array after joining.
	// So to get that first we can do merge of two array then sort them.
	// but here the required time complexity is O(log(m+n)).
	// Whenever we see logy we should first think of binary search.
	// Now here merging of two array will have time complexity of O(m+n)
	// And if we peform binary search on the merged array it will be O(log(m+n)).
	// so we should try merging the array along with binary search.
	//
	// So how do we do it.
	// Let's analyze the array.
	// For example, suppose we have arr1 = [1,2,3,4,5,6,7,8], and arr2 = [3,4,5,6]
	// after joining two arrays it will be arr3 = arr1 + arr2 = [1,2,3,3,4,4,5,5,6,6,7,8]
	// and then median will be 4 at index 6 (since mid = n/2 for odd and (n/2)-1 for even).
	// the left side of array will be l3 = [1,2,3,3,4,4] and r3 = [5,5,6,6,7,8]
	//
	// Here if we look closely we know that to form left side of mid of arr1 and arr2,
	// we need half of arr1 and half of arr2, same the right part will be formed with
	//  right side of arr1 and arr2
	//
	// now we need to check if the left side of arr1 and arr2 is correct and forms
	// the left of merged array. For that we know that the last value of left side
	// of arr1 will be less than the first value of right side of arr1. And for arr2
	// it will follow the same because both arrays are sorted. Which means to get the correct
	// left of merged array we need to have that last value of arr1 <= first value of right
	// side of arr2 && the last value of left side of arr2 <= the first value of right
	// side of arr1. which is,
	//
	// arr1[mid1] <= arr2[mid2+1] && arr2[mid2] <= arr1[mid+1]
	//
	// If that's the case we know that we have found our left of the array, l3 = arr1[mid1]+arr2[mid2]
	// and right of the array, r3 = arr1[mid1+1]+arr2[mid2+1]
	//
	// Now to get the mid value we need to take (max(arr1[mid1],arr2[mid2]) + min(arr1[mid1+1],arr2[mid2+1])) // 2
	// if the length of final array is even, else we will return the last value of l3.
	// But there is a corner case, so for eg- suppose we have an arr1 = [1,2,3,4,5,6,7,8] and
	// arr2 = [1,2,3,4] after using above procedure we get l1 = [1,2,3,4] and l2 = [1,2] and
	// r1 = [5,6,7,8] and r2 = [3,4].here we try to verify if the chosen left side of
	// arrays are correct. so,
	//
	// l1[last index] <= r2[first index] = false and l2[last index] <= r1[first index] = true.
	//
	// so here we can see that the selected values for forming left side of merged array is not
	// correct. To make this correct, we need lose one value from right most of l1 to r1 and gain one
	// value on l2 from r2. now this makes it,
	// l1 = [1,2,3], l2 = [1,2,3], r1 = [4,5,6,7,8], r2 = [4], let's check it again for correct parts
	//
	// l1[last index] <= r2[first index] = true and l2[last index] <= r1[first index] = true
	//
	// This means the partition is correct.
	// There is another corner case for eg, if we don't take any value for one of the array.
	// To cover that we need to use -ve infinity and +ve infinity.
	// we we don't take any value from arr1 then set it to -ve infinity and the full value taken
	// then it will be +ve infinity

	return 0

}
