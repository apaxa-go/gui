// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

CGColorRef createColor(CGFloat* color) {
	CGColorSpaceRef colorSpace = CGColorSpaceCreateWithName(kCGColorSpaceSRGB);
	CGColorRef      r          = CGColorCreate(colorSpace, color);
	CFRelease(colorSpace);
	return r;
}

// color may by nil (at least for GetTextLineGeometry which currently does not use color).
CFAttributedStringRef createAttrStringRef(CFStringRef str, CTFontRef font, CGFloat* color) {
	CFStringRef keys[]   = {kCTFontAttributeName, kCTForegroundColorAttributeName};
	CFTypeRef   values[] = {font, color == nil ? nil : createColor(color)};
	//CFRelease(c); // TODO

	CFDictionaryRef attributes = CFDictionaryCreate(
	    kCFAllocatorDefault, //
	    (const void**)&keys,
	    (const void**)&values,
	    (color == nil ? sizeof(keys) - 1 : sizeof(keys)) / sizeof(keys[0]),
	    &kCFTypeDictionaryKeyCallBacks,
	    &kCFTypeDictionaryValueCallBacks);

	CFAttributedStringRef attrString = CFAttributedStringCreate(kCFAllocatorDefault, str, attributes);
	CFRelease(attributes);
	return attrString;
}

CGSize GetTextImageGeometry(CGContextRef context, CFStringRef str, CTFontRef font) {
	CFAttributedStringRef attrStr = createAttrStringRef(str, font, nil);
	CTLineRef             line    = CTLineCreateWithAttributedString(attrStr);
	CFRelease(attrStr);
	CGSize r = CTLineGetImageBounds(line, context).size;
	/*NSLog(@"Size: %f, Sum: %f, Ascent: %f, Descent: %f, Leading: %f, Box: %@, Bounds: %@",
		CTFontGetSize(font),
		CTFontGetAscent(font)+CTFontGetDescent(font)+CTFontGetLeading(font),
		CTFontGetAscent(font),
		CTFontGetDescent(font),
		CTFontGetLeading(font),
		CGRectCreateDictionaryRepresentation(CTFontGetBoundingBox(font)),
		CGRectCreateDictionaryRepresentation(r)
	);*/
	CFRelease(line);
	return r;
}

void DrawTextImage(CGContextRef context, CFStringRef str, CTFontRef font, CGFloat* color, CGPoint pos) {
	CFAttributedStringRef attrStr = createAttrStringRef(str, font, color);
	CTLineRef             line    = CTLineCreateWithAttributedString(attrStr);
	CFRelease(attrStr);
	CGRect  bounds  = CTLineGetImageBounds(line, context);
	CGPoint prevPos = CGContextGetTextPosition(context);
	pos.x += +prevPos.x - bounds.origin.x;
	pos.y += -prevPos.y + bounds.origin.y + bounds.size.height;
	CGContextSetTextPosition(context, pos.x, pos.y);
	CGContextSaveGState(context);
	CTLineDraw(line, context);
	CFRelease(line);
	CGContextRestoreGState(context);
}

struct TextLineGeometry {
	CGFloat width;
	CGFloat ascent;
	CGFloat descent;
	CGFloat leading;
};

struct TextLineGeometry
    GetTextLineGeometry(CGContextRef context, CFStringRef str, CTFontRef font) {
	CFAttributedStringRef attrStr = createAttrStringRef(str, font, nil);
	CTLineRef             line    = CTLineCreateWithAttributedString(attrStr);
	CFRelease(attrStr);
	struct TextLineGeometry r;
	r.width = CTLineGetTypographicBounds(line, &r.ascent, &r.descent, &r.leading);
	/*NSLog(@"Size: %f, Sum: %f, Ascent: %f, Descent: %f, Leading: %f",
		CTFontGetSize(font),
		r.ascent+r.descent+r.leading,
		r.ascent,
		r.descent,
		r.leading
	);*/
	CFRelease(line);
	return r;
}

void DrawTextLine(CGContextRef context, CFStringRef str, CTFontRef font, CGFloat* color, CGPoint pos, uint8_t origin) {
	CFAttributedStringRef attrStr = createAttrStringRef(str, font, color);
	CTLineRef             line    = CTLineCreateWithAttributedString(attrStr);
	CFRelease(attrStr);
	switch (origin) {
	case 0: pos.y += CTFontGetAscent(font); break;
	case 2: pos.y -= CTFontGetDescent(font); break;
	}
	CGContextSetTextPosition(context, pos.x, pos.y);
	CGContextSaveGState(context);
	CTLineDraw(line, context);
	CFRelease(line);
	CGContextRestoreGState(context);
}