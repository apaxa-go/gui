// Copyright © 2018 Anton Bekker. All rights reserved.
// This file is part of the AGui.
// License information can be found in the LICENSE file.

#import <CoreFoundation/CoreFoundation.h>
#import <CoreGraphics/CoreGraphics.h>
#import <CoreText/CoreText.h>

void resetTransform(CGContextRef context) {
	// TODO is this valid?
	// TODO this implementation vs TopView implementation
	CGContextConcatCTM(context, CGContextGetUserSpaceToDeviceSpaceTransform(context));
}

void fillRoundedRectangle(CGContextRef context, CGRect rect, CGFloat radiusX, CGFloat radiusY, CGFloat* color) {
	CGContextSetFillColor(context, color);
	CGPathRef path = CGPathCreateWithRoundedRect(rect, radiusX, radiusY, NULL);
	CGContextAddPath(context, path);
	CGContextFillPath(context);
	CGPathRelease(path);
}

void drawRoundedRectangle(CGContextRef context, CGRect rect, CGFloat radiusX, CGFloat radiusY, CGFloat* color, CGFloat width) {
	CGContextSetStrokeColor(context, color);
	CGContextSetLineWidth(context, width);
	CGPathRef path = CGPathCreateWithRoundedRect(rect, radiusX, radiusY, NULL);
	CGContextAddPath(context, path);
	CGContextStrokePath(context);
	CGPathRelease(path);
}

// Add 1/4 part of ellipse from current point to (x,y) clockwise.
// Ellipse is horizontal or vertical.
void CGPathEllipseTo(CGMutablePathRef path, CGFloat x, CGFloat y) {
	const CGFloat magic = 4 * (sqrt(2) - 1) / 3; // Well known 0.5522847...

	CGPoint p0 = CGPathGetCurrentPoint(path);

	CGFloat tmpx = (x - p0.x) * (1 - magic);
	CGFloat tmpy = (y - p0.y) * (1 - magic);

	CGFloat cp1x;
	CGFloat cp1y;
	CGFloat cp2x;
	CGFloat cp2y;

	if ((x >= p0.x && y >= p0.y) || (x <= p0.x && y <= p0.y)) { // TODO avoid "if"?
		cp1x = p0.x + (x - p0.x) * magic;
		cp1y = p0.y;
		cp2x = x;
		cp2y = y - (y - p0.y) * magic;
	} else {
		cp1x = p0.x;
		cp1y = p0.y + (y - p0.y) * magic;
		cp2x = x - (x - p0.x) * magic;
		cp2y = y;
	}

	CGPathAddCurveToPoint(path, NULL, cp1x, cp1y, cp2x, cp2y, x, y);
}

// Valid for "flipped" coordinate system (Y-axis from top to bottom; Y=0 is top)
CGMutablePathRef CGPathCreateWithRoundedRectExtended(CGRect rect, CGPoint radiusLT, CGPoint radiusRT, CGPoint radiusRB, CGPoint radiusLB) {
	CGMutablePathRef path = CGPathCreateMutable();
	/*
          1------2
        8          3
        |          |
        |          |
        7          4
          6------5
    */
	CGFloat x1 = rect.origin.x + radiusLT.x;
	CGFloat y1 = rect.origin.y;
	CGFloat x2 = rect.origin.x + rect.size.width - radiusRT.x;
	CGFloat y2 = rect.origin.y;
	CGFloat x3 = rect.origin.x + rect.size.width;
	CGFloat y3 = rect.origin.y + radiusRT.y;
	CGFloat x4 = rect.origin.x + rect.size.width;
	CGFloat y4 = rect.origin.y + rect.size.height - radiusRB.y;
	CGFloat x5 = rect.origin.x + rect.size.width - radiusRB.x;
	CGFloat y5 = rect.origin.y + rect.size.height;
	CGFloat x6 = rect.origin.x + radiusLB.x;
	CGFloat y6 = rect.origin.y + rect.size.height;
	CGFloat x7 = rect.origin.x;
	CGFloat y7 = rect.origin.y + rect.size.height - radiusLB.y;
	CGFloat x8 = rect.origin.x;
	CGFloat y8 = rect.origin.y + radiusLT.y;

	CGPathMoveToPoint(path, NULL, x1, y1);
	CGPathAddLineToPoint(path, NULL, x2, y2);
	CGPathEllipseTo(path, x3, y3);
	CGPathAddLineToPoint(path, NULL, x4, y4);
	CGPathEllipseTo(path, x5, y5);
	CGPathAddLineToPoint(path, NULL, x6, y6);
	CGPathEllipseTo(path, x7, y7);
	CGPathAddLineToPoint(path, NULL, x8, y8);
	CGPathEllipseTo(path, x1, y1);
	CGPathCloseSubpath(path);

	return path;
}

void fillRoundedRectangleExtended(
    CGContextRef context,
    CGRect       rect,
    CGPoint      radiusLT,
    CGPoint      radiusRT,
    CGPoint      radiusRB,
    CGPoint      radiusLB,
    CGFloat*     color //
) {
	CGContextSetFillColor(context, color);
	CGPathRef path = CGPathCreateWithRoundedRectExtended(rect, radiusLT, radiusRT, radiusRB, radiusLB);
	CGContextAddPath(context, path);
	CGContextFillPath(context);
	CGPathRelease(path);
}

void drawRoundedRectangleExtended(
    CGContextRef context,
    CGRect       rect,
    CGPoint      radiusLT,
    CGPoint      radiusRT,
    CGPoint      radiusRB,
    CGPoint      radiusLB,
    CGFloat*     color,
    CGFloat      width //
) {
	CGContextSetStrokeColor(context, color);
	CGContextSetLineWidth(context, width);
	CGPathRef path = CGPathCreateWithRoundedRectExtended(rect, radiusLT, radiusRT, radiusRB, radiusLB);
	CGContextAddPath(context, path);
	CGContextStrokePath(context);
	CGPathRelease(path);
}

void drawConnectedLines(CGContextRef context, CGPoint* points, size_t count) {
	if (count < 2) { return; }

	CGMutablePathRef path = CGPathCreateMutable();
	CGPathAddLines(path, NULL, points, count);
	CGContextAddPath(context, path);
	CGContextStrokePath(context);
	CGPathRelease(path);
}

void fillStraightContour(CGContextRef context, CGPoint* points, size_t count) {
	if (count < 3) { return; }

	CGMutablePathRef path = CGPathCreateMutable();
	CGPathAddLines(path, NULL, points, count);
	CGPathCloseSubpath(path);
	CGContextAddPath(context, path);
	CGContextFillPath(context);
	CGPathRelease(path);
}

void drawStraightContour(CGContextRef context, CGPoint* points, size_t count) {
	if (count < 3) { return; }

	CGMutablePathRef path = CGPathCreateMutable();
	CGPathAddLines(path, NULL, points, count);
	CGPathCloseSubpath(path);
	CGContextAddPath(context, path);
	CGContextStrokePath(context);
	CGPathRelease(path);
}