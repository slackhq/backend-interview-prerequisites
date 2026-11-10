// swift-tools-version: 5.9
import PackageDescription

let package = Package(
    name: "InterviewSetup",
    targets: [
        .systemLibrary(name: "CSQLite", pkgConfig: "sqlite3",
                       providers: [.apt(["libsqlite3-dev"]), .brew(["sqlite3"])]),
        .executableTarget(name: "Setup", dependencies: ["CSQLite"]),
    ],
    swiftLanguageVersions: [.v5]
)
