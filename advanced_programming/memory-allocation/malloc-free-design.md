# Design of Malloc and Free

## Functional Goals
* Malloc allocates memory
* Free deallocates the memory segment allocated by Malloc

## Rules
* You can allocate as much memory as `mmap` allows.
* You can only free an allocated memory segment.

## Performance Goals
* Memory allocation and deallocation should be fast
* Reduce memory fragments
* `mmap` is expensive, don't call it often

## Allocation Philosophy
* If you have a group of free and non-contiguous memory segments, how do you allocate memory? Do you allocate based on the first-fit or do you search for the best fit?
  * First fit is usually faster but might increase fragmentation
  * Best fit is usually slower but might reduces fragmentation
* How do you keep track of currently allocated regions?
* How do you keep track of free memory locations?

## Basic Memory Data Structure
* What data structure will you use to keep track of memory? A sorted list? a tree? a binary heap? a linkedlist?
  * You can easily deframent when you use a sorted list, binary heap or tree but  not with a linkedlist

## Actual Design
* Malloc(size int):
  * Parameters:
    * const int requestSize
    * sortedList freeBlocks
    * sortedList allocatedBlocks
  * When malloc is first called, we check if there are free blocks.
    * If there are no free blocks, we request memory of size requestSize using `mmap` and place in freeBlocks.
    * If there are free blocks available, we use bestFit algorithm to find a free block for the request.
    * If there are no blocks that fit the request, we request memory of size requestSize using `mmap` and place in freeBlocks.
    * If the request to `mmap` fails, throw error
  * How does our bestFit algorithm work?
    * We iterate through the sorted freeBlocks structure and choose the smallest free block that can accomodate the malloc request.
    * With this  algorithm, the performance of malloc is linear to the freeBlocks structure
  * Note that we only allocate from the front of a block!

* Free(name):
  * Get the address of the variable `name`
  * Check if the address is in the list of allocated memory
    * If it is in the list, you can simply move that block to the sortedList freeBlocks
      * If the inserted block is a continuation of an already existing block in freeBlocks, then coalesce them.
    * If it is not in the list, throw an error.

* Fragmentation:
  * We reduce fragmentation a little when we free because we coalesce nearby blocks.
  * We also reduce fragmentation when we allocate from the head of a block.
  * However, that might not be sufficient but if we try to defragment, that would require memory copy which is expensive as well.
  * Is the space gain of deframenting worth the performance hit of memory copy?
