#include <bits/stdc++.h>

int main() {
    std::vector<int> v;

    for(int i=0; i<10; i++) {
        v.push_back(i);
    }

    std::cout << "The elements in the vector: ";
    for(auto it=v.begin(); it != v.end(); it++) {
        std::cout << *it << " ";
    }
    std::cout << "\nFront element: " << v.front() << std::endl;
    std::cout << "Last element: " << v.back() << std::endl;
    std::cout << "size: " << v.size() << std::endl;
    std::cout << "Deleting element from end: " << v[v.size()-1] << std::endl;

    v.pop_back();

    std::cout << "printing after removing: ";
    for(auto it=v.begin(); it!=v.end(); it++) {
        std::cout << *it << " ";
    }
    std::cout << std::endl;

    std::cout << "inserting 5 at the beginning" << std::endl;
    v.insert(v.begin(), 5);

    std::cout << "the first element: " << v[0] << std::endl;

    std::cout << "Erasing the first element" << std::endl;
    v.erase(v.begin());
    std::cout << "first element: " << v[0] << std::endl;

    if(v.empty()) {
        std::cout << "vector is empty" << std::endl;
    }else {
        std::cout << "vector is not empty" << std::endl;
    }

    v.clear();
    std::cout << "size of vector after clearing: " << v.size() << std::endl;

    return 0;
}