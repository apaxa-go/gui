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

void DrawTextLine(CGContextRef context, CFStringRef str, CTFontRef font, CGFloat* color, CGPoint pos) {
	CFAttributedStringRef attrStr = createAttrStringRef(str, font, color);
	CTLineRef             line    = CTLineCreateWithAttributedString(attrStr);
	CFRelease(attrStr);
	CGContextSetTextPosition(context, pos.x, pos.y);
	CGContextSaveGState(context);
	CTLineDraw(line, context);
	CFRelease(line);
	CGContextRestoreGState(context);
}

CGRect GetTextLineGeometry(CGContextRef context, CFStringRef str, CTFontRef font) {
	CFAttributedStringRef attrStr = createAttrStringRef(str, font, nil);
	CTLineRef             line    = CTLineCreateWithAttributedString(attrStr);
	CFRelease(attrStr);
	CGRect r = CTLineGetImageBounds(line, context);
	CFRelease(line);
	return r;
}