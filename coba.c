#include <windows.h>
#include <stdio.h>

typedef int (*add_func)(int, int);
typedef char* (*greet_func)(char*);

int main() {
    HINSTANCE hinstLib;
    add_func add;
    greet_func gret;
    int result;
    char* res2;

    // Load the DLL.
    hinstLib = LoadLibrary("utils.dll");
    if (hinstLib == NULL) {
        printf("Failed to load mydll.dll\n");
        return 1;
    }

    // Get a pointer to the add function in the DLL.
    add = (add_func)GetProcAddress(hinstLib, "SumUtil");
    if (add == NULL) {
        printf("Failed to find add function in mydll.dll\n");
        FreeLibrary(hinstLib);
        return 1;
    }

    gret = (greet_func)GetProcAddress(hinstLib, "Greeting");
    if (gret == NULL) {
        printf("error print");
        FreeLibrary(hinstLib);
        return 1;
    }

    // Call the add function in the DLL.
    result = add(3, 4);
    printf("The result is %d\n", result);

    res2 = gret("sam");
    printf("coba is %s\n", res2);


    // Unload the DLL.
    FreeLibrary(hinstLib);

    return 0;
}
