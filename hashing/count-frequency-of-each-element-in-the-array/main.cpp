#include <bits/stdc++.h>

int main() {
    int arr[] = {10, 5, 10, 15, 10, 5};

    std::map<int, int> mp;
    for(int i=0; i<(sizeof(arr)/sizeof(arr[0])); i++) {
        mp[arr[i]]++;
    }

    for(auto x : mp) {
        std::cout << x.first << " " << x.second << std::endl;
    }

    return 0;
}