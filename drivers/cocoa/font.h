#import <CoreText/CoreText.h>
#include "common.h"

const CGAffineTransform flipped = (CGAffineTransform){1,0,0,1,0,0};//{1,0,0,-1,0,0};

// Translate font-spec italic ([0;1], non-italic 0, italic 1) to core text boolean italic flag.
bool translateItalic(CGFloat i){
    return i>=0.5;
}

// Translate font-spec slant ((-90;+90), upright 0) to core text slant ([-1;1], upright 0, -1 = -30 degrees, 1 = 30 degrees)
CGFloat translateSlant(CGFloat s){
    if (s<= -30){
        return -1;
    }
    if (s>=30){
        return 1;
    }
    // [-30;+30] => [-1;+1]
    return s/30;
}

// Translate font-spec width ((0;+inf), normal 100) to core text weight ([-1;1], normal 0)
CGFloat translateWidth(CGFloat w){
    if (w>=200){
        return 1;
    }
    // (0;100] => [-1;0]
    return (w-100)/100;
}

// Translate font-spec weight ([1;1000], normal 400) to core text weight ([-1;1], normal 0)
CGFloat translateWeight(CGFloat w){
    if (w<400){
        // [1;400] => [-1;0]
        return (w-400)/399;
    }else{
        // [400;1000] => [0;1]
        return (w-400)/600;
    }
}

CTFontDescriptorRef makeFontDescriptor(
    bool reqName,
    UInt8* name,
    CFIndex nameLen,

    bool reqFamily,
    UInt8* family,
    CFIndex familyLen,

    CGFloat size,

    bool reqMonospace,
    bool monospace,

    bool reqItalic,
    CGFloat italic,

    bool reqSlant,
    CGFloat slant,

    bool reqWidth,
    CGFloat width,

    bool reqWeight,
    CGFloat weight
){
    CFMutableDictionaryRef a = CFDictionaryCreateMutable(NULL,0,NULL,NULL);// [NSMutableDictionary dictionary]; // attributes // TODO may return NULL
    CFMutableDictionaryRef t = CFDictionaryCreateMutable(NULL,0,NULL,NULL);//[NSMutableDictionary dictionary]; // traits // TODO may return NULL
    uint32_t st; // symbolic traits // TODO may be set kCTFontUIOptimizedTrait ?

    if (reqName){
        CFStringRef tmp = makeStringRef(name,nameLen);
        CFDictionarySetValue(a,kCTFontNameAttribute,tmp);//[a setObject:tmp forKey:(id)kCTFontNameAttribute];
        CFRelease(tmp);
    }

    if (reqFamily){
        CFStringRef tmp = makeStringRef(family,familyLen);
        CFDictionarySetValue(a,kCTFontFamilyNameAttribute,tmp); //[a setObject:tmp forKey:(id)kCTFontFamilyNameAttribute];
        CFRelease(tmp);
    }

    {
        CFNumberRef tmp = CFNumberCreate(NULL, kCFNumberFloat64Type, &size);
        CFDictionarySetValue(a,kCTFontSizeAttribute,tmp); //[a setObject:tmp forKey:(id)kCTFontSizeAttribute];
        CFRelease(tmp);
    }

    if (reqMonospace){
        if (monospace){
            st|=kCTFontMonoSpaceTrait;
        }
    }

    if (reqItalic){
        if (translateItalic(italic)){
            st|=kCTFontItalicTrait;
        }else{
            // TODO how to announce what we require non-italic font?
        }
    }

    if (reqSlant){
        slant=translateSlant(slant);
        CFNumberRef tmp = CFNumberCreate(NULL, kCFNumberFloat64Type, &slant);
        CFDictionarySetValue(t,kCTFontSlantTrait,tmp); //[t setObject:tmp forKey:(id)kCTFontSlantTrait];
        CFRelease(tmp);
    }

    if (reqWidth){
        width=translateWidth(width);
        CFNumberRef tmp = CFNumberCreate(NULL, kCFNumberFloat64Type, &width);
        CFDictionarySetValue(t,kCTFontWidthTrait,tmp);// [t setObject:tmp forKey:(id)kCTFontWidthTrait];
        CFRelease(tmp);
    }

    if (reqWeight){
        weight=translateWeight(weight);
        CFNumberRef tmp = CFNumberCreate(NULL, kCFNumberFloat64Type, &weight);
        CFDictionarySetValue(t,kCTFontWeightTrait,tmp);//[t setObject:tmp forKey:(id)kCTFontWeightTrait];
        CFRelease(tmp);
    }

    if (st!=0){
        CFNumberRef tmp = CFNumberCreate(NULL, kCFNumberSInt32Type, &st);
        CFDictionarySetValue(t,kCTFontSymbolicTrait,tmp);//[t setObject:tmp forKey:(id)kCTFontSymbolicTrait];
        CFRelease(tmp);
    }

    CFDictionarySetValue(a,kCTFontTraitsAttribute,t);//[a setObject:t forKey:(id)kCTFontTraitsAttribute];
    CFRelease(t);
    return CTFontDescriptorCreateWithAttributes(a);
}

CTFontRef makeDefaultFont(
    CGFloat size,

    bool reqMonospace,
    bool monospace,

    bool reqItalic,
    CGFloat italic,

    bool reqSlant,
    CGFloat slant,

    bool reqWidth,
    CGFloat width,

    bool reqWeight,
    CGFloat weight
){
    CTFontRef tmp = CTFontCreateUIFontForLanguage(kCTFontUIFontUser, size, NULL);
    CTFontDescriptorRef descriptor = makeFontDescriptor(
        false,
        NULL,
        0,

        false,
        NULL,
        0,

        size,

        reqMonospace,
        monospace,

        reqItalic,
        italic,

        reqSlant,
        slant,

        reqWidth,
        width,

        reqWeight,
        weight
    );
    CTFontRef f = CTFontCreateCopyWithAttributes(tmp, size, &flipped, descriptor);
    CFRelease(tmp);
    return f;
}

CTFontRef makeFont(
    bool reqName,
    UInt8* name,
    CFIndex nameLen,

    bool reqFamily,
    UInt8* family,
    CFIndex familyLen,

    CGFloat size,

    bool reqMonospace,
    bool monospace,

    bool reqItalic,
    CGFloat italic,

    bool reqSlant,
    CGFloat slant,

    bool reqWidth,
    CGFloat width,

    bool reqWeight,
    CGFloat weight
){
    CTFontDescriptorRef descriptor = makeFontDescriptor(
        reqName,
        name,
        nameLen,

        reqFamily,
        family,
        familyLen,

        size,

        reqMonospace,
        monospace,

        reqItalic,
        italic,

        reqSlant,
        slant,

        reqWidth,
        width,

        reqWeight,
        weight
    );

    if (descriptor==NULL) {
        return NULL;
    }

    CTFontRef font = CTFontCreateWithFontDescriptor(descriptor, size, &flipped);
    CFRelease(descriptor);

    return font;
}

CTFontRef makeFontFromFile(
    UInt8* path,
    CFIndex pathLen,

    CGFloat size,

    bool reqMonospace,
    bool monospace,

    bool reqItalic,
    CGFloat italic,

    bool reqSlant,
    CGFloat slant,

    bool reqWidth,
    CGFloat width,

    bool reqWeight,
    CGFloat weight
){
    // TODO implement index access for collection
    CTFontRef CTFont;

    CFStringRef pathRef = makeStringRef(path,pathLen);
    CFURLRef url = CFURLCreateWithFileSystemPath(NULL, pathRef, kCFURLPOSIXPathStyle, false);
    CGDataProviderRef dataProvider = CGDataProviderCreateWithURL(url);
    if (dataProvider!=NULL){
        CGFontRef CGFont = CGFontCreateWithDataProvider(dataProvider);
        if (CGFont!=NULL){
            CTFontDescriptorRef descriptor = makeFontDescriptor(
                false,
                NULL,
                0,

                false,
                NULL,
                0,

                size,

                reqMonospace,
                monospace,

                reqItalic,
                italic,

                reqSlant,
                slant,

                reqWidth,
                width,

                reqWeight,
                weight
            );
            CTFont = CTFontCreateWithGraphicsFont(CGFont,size,&flipped,descriptor);
            CFRelease(descriptor);
            CFRelease(CGFont);
        }
        CFRelease(dataProvider);
    }
    CFRelease(url);
    CFRelease(pathRef);

    return CTFont;
}

void releaseFont(CTFontRef font){
    CFRelease(font);
}