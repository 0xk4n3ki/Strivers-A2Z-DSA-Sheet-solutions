#include <bits/stdc++.h>

int main() {
    std::vector<int> v = {4, 2, 5, 9, 1};
    std::cout << "The element in the vector are: ";

    for(int i=0; i<v.size(); i++) {
        std::cout << v[i] << " ";
    }
    std::cout << std::endl;

    std::cout << "The maximum element is: " << *std::max_element(v.begin(), v.end()) << std::endl;

    return 0;
}