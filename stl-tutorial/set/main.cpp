#include <bits/stdc++.h>

int main() {
    std::set<int> s;
    for(int i=0; i<=10; i++) {
        s.insert(i);
    }

    std::cout << "Elements present in the set: ";
    for(auto it=s.begin(); it!=s.end(); it++) {
        std::cout << *it << " ";
    }
    std::cout << std::endl;

    int n=2;
    if(s.find(2) != s.end()) {
        std::cout << n << " is present in set" << std::endl;
    }

    s.erase(s.begin());
    std::cout << "Elements after deleting first element: ";
    for(auto it=s.begin(); it!=s.end(); it++) {
        std::cout << *it << " ";
    }
    std::cout << std::endl;

    std::cout << "Size of the set: " << s.size() << std::endl;

    if(s.empty() == false) {
        std::cout << "The set is not empty" << std::endl;
    }else {
        std::cout << "the set is empty" << std::endl;
    }
    s.clear();

    std::cout << "size after clearing: " << s.size() << std::endl;

    return 0;
}