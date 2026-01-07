#include <bits/stdc++.h>

int main() {
    std::unordered_multiset<int> s;
    for(int i=0; i<=10; i++) {
        s.insert(i);
    }

    s.insert(5);

    std::cout << "elements present in the unordered multiset: ";
    for(auto it = s.begin(); it!=s.end(); it++) {
        std::cout << *it << " ";        
    }
    std::cout << std::endl;

    int n=2;
    if(s.find(2) != s.end()) {
        std::cout << n << " is present in unordered multiset" << std::endl;
    }
    s.erase(s.begin());
    s.erase(5);
    std::cout << "Elements after deleting the first element: ";
    for(auto it=s.begin(); it!=s.end(); it++) {
        std::cout << *it << " ";
    }
    std::cout << std::endl;

    std::cout << "the size of the unordered multiset is: " << s.size() << std::endl;
    
    if(s.empty() == false) {
        std::cout << "the unordered multiset is not empty" << std::endl;
    }else {
        std::cout << "the unordered multiset is empty" << std::endl;
    }
    s.clear();
    std::cout << "size of the unordered multiset after clearning all the elements: " << s.size() << std::endl;

    return 0;
}