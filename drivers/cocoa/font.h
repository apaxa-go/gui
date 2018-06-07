// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

#import <CoreText/CoreText.h>

//const CGAffineTransform flipped = (CGAffineTransform){1,0,0,1,0,0};//{1,0,0,-1,0,0};

struct FontSpec {
	CGFloat size;
	bool    reqMonospace;
	bool    monospace;
	bool    reqItalic;
	CGFloat italic;
	bool    reqSlant;
	CGFloat slant;
	bool    reqWidth;
	CGFloat width;
	bool    reqWeight;
	CGFloat weight;
};

// Translate font-spec italic ([0;1], non-italic 0, italic 1) to core text boolean italic flag.
bool translateItalic(CGFloat i) { return i >= 0.5; }

// Translate font-spec slant ((-90;+90), upright 0) to core text slant ([-1;1], upright 0, -1 = -30 degrees, 1 = 30 degrees)
CGFloat translateSlant(CGFloat s) {
	if (s <= -30) { return -1; }
	if (s >= 30) { return 1; }

	// [-30;+30] => [-1;+1]
	return s / 30;
}

// Translate font-spec width ((0;+inf), normal 100) to core text weight ([-1;1], normal 0)
CGFloat translateWidth(CGFloat w) {
	if (w >= 200) { return 1; }

	// (0;100] => [-1;0]
	return (w - 100) / 100;
}

// Translate font-spec weight ([1;1000], normal 400) to core text weight ([-1;1], normal 0)
CGFloat translateWeight(CGFloat w) {
	if (w < 400) {
		// [1;400] => [-1;0]
		return (w - 400) / 399;
	} else {
		// [400;1000] => [0;1]
		return (w - 400) / 600;
	}
}

// name & family may be NULL.
CTFontDescriptorRef CreateFontDescriptor(CFStringRef name, CFStringRef family, struct FontSpec spec) {
	CFMutableDictionaryRef a = CFDictionaryCreateMutable(NULL, 0, NULL, NULL); // attributes, in some cases may returns NULL
	CFMutableDictionaryRef t = CFDictionaryCreateMutable(NULL, 0, NULL, NULL); // traits, in some cases may returns NULL
	uint32_t               st = 0; // symbolic traits // TODO may be set kCTFontUIOptimizedTrait ?

	if (name != nil) { CFDictionarySetValue(a, kCTFontNameAttribute, name); }
	if (family != nil) { CFDictionarySetValue(a, kCTFontFamilyNameAttribute, family); }
	{
		CFNumberRef tmp = CFNumberCreate(NULL, kCFNumberFloat64Type, &spec.size);
		CFDictionarySetValue(a, kCTFontSizeAttribute, tmp);
		CFRelease(tmp);
	}
	if (spec.reqMonospace) {
		if (spec.monospace) { st |= kCTFontMonoSpaceTrait; }
	}
	if (spec.reqItalic) {
		if (translateItalic(spec.italic)) {
			st |= kCTFontItalicTrait;
		} else {
			// TODO how to announce what we require non-italic font?
		}
	}
	if (spec.reqSlant) {
		spec.slant      = translateSlant(spec.slant);
		CFNumberRef tmp = CFNumberCreate(NULL, kCFNumberFloat64Type, &spec.slant);
		CFDictionarySetValue(t, kCTFontSlantTrait, tmp);
		CFRelease(tmp);
	}
	if (spec.reqWidth) {
		spec.width      = translateWidth(spec.width);
		CFNumberRef tmp = CFNumberCreate(NULL, kCFNumberFloat64Type, &spec.width);
		CFDictionarySetValue(t, kCTFontWidthTrait, tmp);
		CFRelease(tmp);
	}
	if (spec.reqWeight) {
		spec.weight     = translateWeight(spec.weight);
		CFNumberRef tmp = CFNumberCreate(NULL, kCFNumberFloat64Type, &spec.weight);
		CFDictionarySetValue(t, kCTFontWeightTrait, tmp);
		CFRelease(tmp);
	}
	if (st != 0) {
		CFNumberRef tmp = CFNumberCreate(NULL, kCFNumberSInt32Type, &st);
		CFDictionarySetValue(t, kCTFontSymbolicTrait, tmp);
		CFRelease(tmp);
	}

	CFDictionarySetValue(a, kCTFontTraitsAttribute, t);
	CFRelease(t);
	CTFontDescriptorRef r = CTFontDescriptorCreateWithAttributes(a);
	CFRelease(a);
	return r;
}

CTFontRef CreateDefaultFont(struct FontSpec spec) {
	CTFontRef           tmp        = CTFontCreateUIFontForLanguage(kCTFontUIFontUser, spec.size, NULL);
	CTFontDescriptorRef descriptor = CreateFontDescriptor(NULL, NULL, spec);
	CTFontRef           f          = CTFontCreateCopyWithAttributes(tmp, spec.size, NULL, descriptor);
	CFRelease(descriptor);
	CFRelease(tmp);
	return f;
}

// name & family may be NULL.
CTFontRef CreateFont(CFStringRef name, CFStringRef family, struct FontSpec spec) {
	CTFontDescriptorRef descriptor = CreateFontDescriptor(name, family, spec);
	if (descriptor == NULL) { return NULL; }
	CTFontRef font = CTFontCreateWithFontDescriptor(descriptor, spec.size, NULL);
	CFRelease(descriptor);
	return font;
}

CTFontRef CreateFontFromFile(CFStringRef path, struct FontSpec spec) {
	// TODO implement index access for collection
	CTFontRef         CTFont       = nil;
	CFURLRef          url          = CFURLCreateWithFileSystemPath(NULL, path, kCFURLPOSIXPathStyle, false);
	CGDataProviderRef dataProvider = CGDataProviderCreateWithURL(url);

	if (dataProvider != NULL) {
		CGFontRef CGFont = CGFontCreateWithDataProvider(dataProvider);
		if (CGFont != NULL) {
			CTFontDescriptorRef descriptor = CreateFontDescriptor(NULL, NULL, spec);
			CTFont                         = CTFontCreateWithGraphicsFont(CGFont, spec.size, NULL, descriptor);
			CFRelease(descriptor);
			CFRelease(CGFont);
		}
		CFRelease(dataProvider);
	}
	CFRelease(url);

	return CTFont;
}
