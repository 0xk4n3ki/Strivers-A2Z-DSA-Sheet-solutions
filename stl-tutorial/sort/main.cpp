#include <bits/stdc++.h>

void printVec(std::vector<int> vec) {
    std::vector<int>::iterator it;
    for(it=vec.begin(); it!=vec.end(); it++) {
        std::cout << *it << " ";
    }
    std::cout << std::endl;
}

bool sortbysec(const std::pair<int, int> &a, const std::pair<int, int> &b) {
    return a.second < b.second;
}

int main() {
    std::vector<int> vec = {4, 3, 1};
    sort(vec.begin(), vec.end());
    printVec(vec);

    sort(vec.begin(), vec.end(), std::greater<int>());
    printVec(vec);

    std::vector<std::pair<int, int>> vec1 = {{10, 3}, {20, 1}, {30, 2}};
    sort(vec1.begin(), vec1.end(), sortbysec);
    for(auto it=vec1.begin(); it!=vec1.end(); it++) {
        std::cout << it->first << " " << it->second << std::endl;
    }
    return 0;
}