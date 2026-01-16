#include <bits/stdc++.h>

void reverseArr(int arr[], int left, int right) {
    if(left >= right) {
        return;
    }
    int tmp = arr[left];
    arr[left] = arr[right];
    arr[right] = tmp;
    reverseArr(arr, left+1, right-1);
}

int main() {
    int n;
    std::cout << "Enter n: ";
    std::cin >> n;

    int arr[n];
    std::cout << "Enter the array numbers: ";
    for(int i=0; i<n; i++) {
        std::cin >> arr[i];
    }

    reverseArr(arr, 0, n-1);
    for(int i=0; i<n; i++) {
        std::cout << arr[i] << " ";
    }
    std::cout << std::endl;
}