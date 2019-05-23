#include<stdio.h>
#include<string.h>

void my_strcpy(char *dst, char *src) {
    int i = 0;

    for(i = 0; src[i] != '\0'; i++){
        dst[i] = src[i];
    }
    dst[i] = 0;

}
void my_strcpy2(char *dst, char *src) {
    int i = 0;

    while (*dst++ = *src++){
        NULL;
    }
    dst[i] = 0;
}
int my_strcpy3(char *dst, char *src) {
    int i = 0;

    while (*dst++ = *src++){
        NULL;
    }
    dst[i] = 0;
}

int main(void) {
    char src[] = "abcdefg";
    char dst[100];

    my_strcpy2(dst,src);
    printf("%s\n",dst);
}
