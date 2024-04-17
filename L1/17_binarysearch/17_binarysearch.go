package binarysearch

func binarySearch(arr []int, left int, right int, searchNum int) int {
    if right >= left {
        var mid = left + (right-left)/2 
        var midVal = arr[mid]

        if midVal == searchNum {
            return mid
        } else if midVal < searchNum {
            return binarySearch(arr, mid+1, right, searchNum) 
        } else {
            return binarySearch(arr, left, mid-1, searchNum) 
        }
    }
    return -1 
}
