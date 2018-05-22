#import <CoreText/CoreText.h>
#include "common.h"

CTFontRef makeFont(const UInt8 *bytes, CFIndex numBytes, CGFloat size){
    CFStringRef name = makeStringRef(bytes, numBytes);
    if (name==NULL) {
        return NULL;
    }

    CTFontDescriptorRef descriptor = CTFontDescriptorCreateWithNameAndSize(name, size);
    if (descriptor==NULL) {
        return NULL;
    }

    CTFontRef font = CTFontCreateWithFontDescriptor(descriptor, 0.0, NULL);
    CFRelease(descriptor);
    CFRelease(name);

    return font;
}

void releaseFont(CTFontRef font){
    CFRelease(font);
}