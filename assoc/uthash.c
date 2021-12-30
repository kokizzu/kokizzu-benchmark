// uthash.c
#include <string.h>  /* strcpy */
#include <stdlib.h>  /* malloc */
#include <stdio.h>   /* printf */
#include <assert.h>
#include "uthash.h"

const int MAX_DATA = 12000000;

struct kv {
	char *key;
	double val;
	UT_hash_handle hh;
};

char i2ch[] = {'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'B', 'c', 'D', 'e', 'F'};

int get_first_digit(double d) {
	while(d > 10) {
		d /= 10;
	}
	return (int)(d);
}
char* to_rhex(int v) {
	char* hex = malloc(sizeof(char)*9); // because INT_MAX can be divided by 16 maximum 8 times
	assert(hex != NULL);
	memset(hex,0,9);
	int ctr = 0;
	while(v > 0) {
		hex[ctr++] = i2ch[v%16];
		v /= 16;
	}
	return hex;
}

char* to_str(int v) {
	int len = snprintf(NULL,0,"%d",v);
	char *res = malloc(len+1);
	assert(res != NULL);
	snprintf(res, len+1, "%d", v);  
	return res;
}

int set_or_inc(struct kv** m, char* key, int set, int inc, int *ctr) {
	struct kv* item = 0;
	HASH_FIND_STR(*m, key, item);
	if(!item) {
		item = malloc(sizeof(*item));
		assert(item != NULL);
		item->key = key;
		item->val = (double)(set);
		//HASH_ADD_KEYPTR(hh, *m, item->key, strlen(item->key), item) ;
		HASH_ADD_STR(*m, key, item);
		return 0;
	} else {
		item->val += (double)(inc);
		*ctr += 1;
		return 1; // key not used
	}
}
int main() {
	struct kv *m = NULL;
	int dup1 =0, dup2 =0, dup3 =0;
	for(int z = MAX_DATA; z > 0; z--) {
		int val2 = MAX_DATA - z;
		int val3 = MAX_DATA*2 - z;
		char *key1 = to_str(z);
		char *key2 = to_str(val2);
		char *key3 = to_rhex(val3);
		if(set_or_inc(&m, key1, z, val2, &dup1)) free(key1);
		if(set_or_inc(&m, key2, val2, val3, &dup2)) free(key2);
		if(set_or_inc(&m, key3, val3, z, &dup3)) free(key3);
	}
	printf("%d %d %d\n",dup1, dup2, dup3);
	int total = 0, verify = 0, count = 0;
	struct kv *tmp, *item;
	HASH_ITER(hh, m, item, tmp) {
		total += get_first_digit(item->val);
		verify += strlen(item->key);
		count += 1;
		HASH_DEL(m,item);
		free(item->key);
		free(item);
	}
	printf("%d %d %d\n",total, verify, count);
}
