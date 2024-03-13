//go:build ignore

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <unistd.h>

#define MAX_LENGTH 100

typedef struct {
    char name[MAX_LENGTH];
    int age;
} Person;

int demonstratePointers(int *value);
char *demonstrateDynamicMemory(const char *input);
Person demonstrateStructs(char *name, int age);
int demonstrateFileIO(const char *message);

int main() {
    srand(time(NULL)); 

    while (1) {
        int choice = rand() % 4; 

        switch (choice) {
            case 0: {
                int x = 10, result;
                result = demonstratePointers(&x);
                printf("Pointer function returned: %d\n", result);
                break;
            }
            case 1: {
                char *result = demonstrateDynamicMemory("Hello, Dynamic World!");
                printf("Dynamic memory function returned: %s\n", result);
                free(result); 
                break;
            }
            case 2: {
                Person p = demonstrateStructs("John Doe", 30);
                printf("Struct function returned: %s, %d\n", p.name, p.age);
                break;
            }
            case 3: {
                int result = demonstrateFileIO("Hello, File World!");
                printf("File IO function returned: %d\n", result);
                break;
            }
        }

        sleep(1); 
    }

    return 0;
}

int demonstratePointers(int *value) {
    printf("Pointer example: %d\n", *value);
    return *value + 10;
}

char *demonstrateDynamicMemory(const char *input) {
    char *str = (char *)malloc(MAX_LENGTH * sizeof(char));
    strcpy(str, input);
    printf("Dynamic memory allocation: %s\n", str);
    return str; 
}

Person demonstrateStructs(char *name, int age) {
    Person person;
    strcpy(person.name, name);
    person.age = age;
    printf("Struct example: %s is %d years old.\n", person.name, person.age);
    return person;
}

int demonstrateFileIO(const char *message) {
    FILE *file = fopen("example.txt", "w");
    if (file == NULL) {
        printf("File opening failed.\n");
        return -1;
    }
    fprintf(file, "%s\n", message);
    fclose(file);

    printf("File I/O example: wrote '%s' to example.txt\n", message);
    return 0; 
}
