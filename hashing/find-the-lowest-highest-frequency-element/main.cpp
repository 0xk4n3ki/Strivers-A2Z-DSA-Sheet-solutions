#include <bits/stdc++.h>

int main() {
    int arr[] = {10, 5, 10, 15, 10, 5};

    std::unordered_map<int, int> mp;
    for(int i=0; i<sizeof(arr)/sizeof(arr[0]); i++) {
        mp[arr[i]]++;
    }

    int min = 10000;
    int minFreq = 100000;
    int max = -1;
    int maxFreq = -1;

    for(auto x : mp) {
        if(x.second < minFreq) {
            minFreq = x.second;
            min = x.first;
        }
        if(x.second > maxFreq) {
            maxFreq = x.second;
            max = x.first;
        }
    }

    std::cout << "min: " << min << ", max: " << max << std::endl;

    return 0;
}
