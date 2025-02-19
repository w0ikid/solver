#include <iostream>

struct Node {
    int data;
    int index;
    Node* next;
};

class LinkedList
{
public:
    Node* head; // указатель на первый элемент списка
    int index_count;

    LinkedList() : head(nullptr), index_count(0) {} // конструктор изначально список пуст

    void append(int newData) {
        Node* newNode = new Node(); // new node
        newNode->data = newData;
        newNode->next = nullptr;
        newNode->index = index_count++;
        if (head == nullptr) {
            head = newNode; // // Если список пуст, новый узел становится первым
        } else {
            Node* last = head;
            while (last->next != nullptr){
                last = last->next; // Находим последний узел
            }
            last->next = newNode;
        }
    }

    void displayAt(int index){
        Node* current = head;

        while (current != nullptr){
            if (current->index == index) {
                std::cout << current->data << std::endl;
                return;
            }
            current = current->next;
        }
    }

    void display() {
        Node* current = head;
        while (current != nullptr) {
            std::cout << current->data << " ";
            current = current->next;
        }
        std::cout << "\n";
    }
    // Деструктор для освобождения памяти
    ~LinkedList() {
        Node* current = head;
        while (current != head) {
            Node* next = current->next;
            delete current;
            current = next;
        }
    }
};


int main(){
    LinkedList list;
    list.append(1);
    list.append(2);
    list.append(3);

    list.displayAt(0);
    list.displayAt(2);
    list.append(15);
    list.display();
}