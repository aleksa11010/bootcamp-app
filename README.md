# Sample Java Application — Task Manager

A simple task manager library demonstrating Java best practices including JUnit 5 testing, parameterized tests, nested test classes, and linting with Checkstyle/PMD/SpotBugs.

**Note:** One test is deliberately failing (`statsReturnsZeroForEmpty`) to demonstrate test failure reporting.

## Project Structure

```
sample-go-app/
├── pom.xml                                          # Maven build config
├── checkstyle.xml                                   # Checkstyle rules
├── README.md
├── .harness/pipeline.yaml                           # Harness CI pipeline
└── src/
    ├── main/java/com/sample/taskmanager/
    │   ├── App.java                                 # CLI entry point
    │   ├── Task.java                                # Task entity
    │   ├── TaskManager.java                         # Manager with CRUD operations
    │   ├── TaskNotFoundException.java               # Custom exception
    │   ├── Priority.java                            # Priority enum
    │   ├── Status.java                              # Status enum
    │   └── Stats.java                               # Aggregate statistics
    └── test/java/com/sample/taskmanager/
        ├── TaskTest.java                            # Tests for Task/Priority/Status
        └── TaskManagerTest.java                     # Tests for TaskManager (1 fails)
```

## Running

```bash
mvn compile exec:java -Dexec.mainClass="com.sample.taskmanager.App"
```

## Testing

```bash
# Run all tests
mvn test

# Run with verbose/debug output
mvn test -X

# Run a specific test class
mvn test -Dtest=TaskManagerTest

# Run a specific test method
mvn test -Dtest="TaskManagerTest#statsReturnsZeroForEmpty"

# JUnit XML reports are automatically generated at:
# target/surefire-reports/TEST-*.xml
```

## Linting

```bash
# Run Checkstyle
mvn checkstyle:check

# Run PMD
mvn pmd:check

# Run SpotBugs (requires compiled classes)
mvn compile spotbugs:check

# Run all linters at once
mvn checkstyle:check pmd:check compile spotbugs:check
```

## Recommended Java Linters

### 1. [Checkstyle](https://checkstyle.org/) — Code Style Enforcement

Enforces naming conventions, formatting, imports, complexity limits. This project includes a `checkstyle.xml` config.

### 2. [PMD](https://pmd.github.io/) — Static Code Analysis

Finds common bugs, dead code, suboptimal code, and overly complex expressions. Configured in `pom.xml`.

### 3. [SpotBugs](https://spotbugs.github.io/) — Bug Detection

Successor to FindBugs. Detects potential bugs via bytecode analysis (null pointer dereferences, infinite loops, etc.).

### 4. [Error Prone](https://errorprone.info/) — Compile-Time Bug Detection

Google's Java compiler plugin that catches common mistakes at compile time.

### 5. [SonarLint](https://www.sonarsource.com/products/sonarlint/) — IDE Integration

Real-time linting in your IDE with 600+ rules for Java.

## Harness CI Commands

Use these commands in your Harness pipeline steps:

| Purpose | Command |
|---|---|
| **Run tests** | `mvn test` |
| **Run tests (JUnit XML output)** | `mvn test` (Surefire generates XML in `target/surefire-reports/`) |
| **Run Checkstyle** | `mvn checkstyle:check` |
| **Run PMD** | `mvn pmd:check` |
| **Run SpotBugs** | `mvn compile spotbugs:check` |
| **Run everything** | `mvn clean verify checkstyle:check pmd:check spotbugs:check` |
