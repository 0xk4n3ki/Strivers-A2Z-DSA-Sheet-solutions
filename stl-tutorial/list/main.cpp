#include <bits/stdc++.h>

void printList(std::list<int> li) {
    std::list<int>::iterator it;
    for(it=li.begin(); it!=li.end(); it++) {
        std::cout << *it << " ";
    }
    std::cout << std::endl;
}

int main() {
    std::list<int> li;
    li.push_back(10);
    li.push_back(20);
    li.push_front(30);
    li.push_front(40);
    li.push_front(50);

    std::cout << "The elements in the list are: ";
    printList(li);

    std::cout << "Reversing the list: ";
    li.reverse();
    printList(li);

    std::cout << "Sorting the list: ";
    li.sort();
    printList(li);

    std::cout << "The size of the list is: " << li.size() << std::endl;
    std::cout << "The first element in the list: " << li.front() << std::endl;
    std::cout << "Deleting the first element" << std::endl;
    li.pop_front();
    printList(li);
    
    std::cout << "The last element of the list: " << li.back() << std::endl;
    std::cout << "Deleting the last element" << std::endl;
    li.pop_back();
    printList(li);

    return 0;
}