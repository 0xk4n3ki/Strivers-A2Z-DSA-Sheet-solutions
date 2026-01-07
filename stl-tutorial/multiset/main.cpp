#include <bits/stdc++.h>

int main() {
    std::multiset<int> s;

    for(int i=1;i<=10;i++) {
        s.insert(i);
    }

    s.insert(5);
    std::cout << "elements present in the multiset: ";
    for(auto it=s.begin(); it!=s.end(); it++) {
        std::cout << *it << " ";
    }
    std::cout << std::endl;

    int n=2;
    if(s.find(2) !=s.end()) {
        std::cout << n << " is present in multiset" << std::endl;
    }

    s.erase(s.begin());
    std::cout << "elements after deleting the first element: ";
    for(auto it=s.begin(); it != s.end(); it++) {
        std::cout << *it << " ";
    }
    std::cout << std::endl;

    std::cout << "the size of the multiset is: " << s.size() << std::endl;
    if(!s.empty())
        std::cout << "multiset is not empty" << std::endl;
    else
        std::cout << "multiset is empty" << std::endl;

    s.clear();
    std::cout << "size of the multiset after clearing: " << s.size() << std::endl;
    return 0;
}