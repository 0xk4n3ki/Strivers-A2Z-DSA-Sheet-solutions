#include <bits/stdc++.h>

int main() {
    std::unordered_set<int> s;

    for(int i=0; i<=10; i++) {
        s.insert(i);
    }
    std::cout << &s << std::endl;

    std::cout << "Elements present in the unordered set: ";
    for(auto it = s.begin(); it != s.end(); it++) {
        std::cout << *it << " ";
    }
    std::cout << std::endl;

    s.erase(s.begin());

    std::cout << "Elements after deleting the arbitrary element: ";
    for(auto it = s.begin(); it != s.end(); it++) {
        std::cout  << *it << " ";
    }
    std::cout << std::endl;

    std::cout << "Size of unordered set: " << s.size() << std::endl;

    if(s.empty() == false)
        std::cout << "unordered set is not empty" << std::endl;
    else
        std::cout << "unordered set is empty" << std::endl;

    s.clear();
    
    std::cout << "Size of the unordered set after clearing all elements: " << s.size() << std::endl;
    return 0;
}