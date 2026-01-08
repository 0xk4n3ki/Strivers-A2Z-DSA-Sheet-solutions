#include <bits/stdc++.h>

void printQueue(std::queue<int> q1) {
    std::queue<int> q2=q1;
    while(!q2.empty()) {
        std::cout << q2.front() << " ";
        q2.pop();
    }
    std::cout << std::endl;
}

int main() {
    std::queue<int> q;
    for(int i=0; i<=5; i++) {
        q.push(i);
    }

    std::cout << "The elements of the queue are: ";
    printQueue(q);

    std::cout << "The size of the queue: " << q.size() << std::endl;
    std::cout << "The fromt element of the queue: " << q.front() << std::endl;
    std::cout << "The last element of the queue: " << q.back() << std::endl;
    std::cout << "Pop the front element: " << std::endl;

    q.pop();
    printQueue(q);

    return 0;
}