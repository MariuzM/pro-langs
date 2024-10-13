const std = @import("std");

pub fn main() void {
    const name = "Hello, World!";
    std.debug.print("{s}", .{name});
}
