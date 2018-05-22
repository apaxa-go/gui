#include "common.h"

CFStringRef makeStringRef(const UInt8 *bytes, CFIndex numBytes){
    // TODO check functions results for nil???
    return CFStringCreateWithBytes(kCFAllocatorDefault, bytes, numBytes, kCFStringEncodingUTF8, false);
}