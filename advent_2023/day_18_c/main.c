#include <stdlib.h>
#include <stdio.h>
#include <string.h>

const char *file = "pre_input.txt";

typedef struct Vec2 {
    int x;
    int y;
} Vec2;

//static Vec2 up(Vec2 vec) {
//    Vec2 v = {.x = vec.x, .y = vec.y - 1};
//    return v;
//}
//
//static Vec2 right(Vec2 vec) {
//    Vec2 v = {.x = vec.x + 1, .y = vec.y};
//    return v;
//}
//
//static Vec2 down(Vec2 vec) {
//    Vec2 v = {.x = vec.x, .y = vec.y + 1};
//    return v;
//}
//
//static Vec2 left(Vec2 vec) {
//    Vec2 v = {.x = vec.x - 1, .y = vec.y};
//    return v;
//}
//
//static Vec2 add(Vec2 vec1, Vec2 vec2) {
//    Vec2 v = {.x = vec1.x + vec2.x, .y = vec1.y + vec2.y};
//    return v;
//}

char *trimFT(int from, int to, char str[]) {
    char *res = malloc(sizeof(char*)*(to-from)+1);
    strncpy(res, &str[from], to-from);
    res[to-from] = '\0'; 
    return res;
}

typedef struct DigInstr {
    Vec2 v;
    int len;
} DigInstr;

DigInstr *readInput(char **lines, int len) {
    //DigIsntr *result = malloc(sizeof(DigIsntr)*len);
    for (int i = 0; i < len; i++) {
        char c[2];
        int a = 0;
        char str[10];
        sscanf(lines[i], "%s %d %s", c, &a, str);
        char *step_hex = trimFT(2, 7, str);
        char *d = trimFT(7,8, str);
        printf("%s | %s\n", step_hex, d);
        unsigned long steps = strtoul(step_hex, NULL, 16);
        printf("%ld\n", steps);
        free(step_hex);
        free(d);
    }
    return NULL;
}

int main(void) {
         FILE * fp;
    char * line = NULL;
    size_t len = 0;
    ssize_t read;

    fp = fopen(file, "r");
    if (fp == NULL) {
        exit(EXIT_FAILURE);
            return 1;
    }

    char **arr = malloc(sizeof(char *)*100);

    int lines_count = 0;
    while ((read = getline(&line, &len, fp)) != -1) {
        char *str = malloc(sizeof(char)*14);
        strncpy(str, line, 14);
        str[13] = '\0';
        arr[lines_count] = str;
        lines_count++;
    }

    DigInstr *dig_isntr = readInput(arr, lines_count);
    if (dig_isntr == NULL) {
        printf("dig_isntr is NULL\n");
    }

    for (int i = 0; i < lines_count; i++) {
        if (arr[i])
            free(arr[i]);
    }

    if (arr)
        free(arr);

    fclose(fp);
    if (line)
        free(line);

    exit(EXIT_SUCCESS);
    return 0;
}
