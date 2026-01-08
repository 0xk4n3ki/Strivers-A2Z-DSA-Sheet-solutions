#include <bits/stdc++.h>

void printDeque(std::deque<int> dq) {
    std::deque<int>::iterator it;
    for(it = dq.begin(); it != dq.end(); it++) {
        std::cout << *it << " ";
    }
    std::cout << std::endl;
}

int main() {
    std::deque<int> dq;
    dq.push_back(10);
    dq.push_back(20);
    dq.push_front(30);
    dq.push_front(40);
    dq.push_front(50);

    std::cout << "Elements in the deque are: ";
    printDeque(dq);

    std::cout << "The size of the deque is: " << dq.size() << std::endl;
    std::cout << "The first element in the deque: " << dq.front() << std::endl;
    std::cout << "Deleting the first element: ";
    dq.pop_front();
    printDeque(dq);

    std::cout << "The last element of the deque: " << dq.back() << std::endl;
    std::cout << "Deleting the last element: ";
    dq.pop_back();
    printDeque(dq);
}