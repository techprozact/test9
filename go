We are simplifying the agent build approach to avoid CI/CD complexity and dependency issues.

Do NOT build the Java agent inside Docker or Tekton pipeline.

Instead:

1. Build perfsonic-agent.jar once during development

2. Commit the generated JAR into:
   frontend/public/perfsonic-agent.jar

3. Also commit:
   frontend/public/run-agent.bat

4. Remove:

   * Maven build step from Dockerfile
   * Node/Maven multi-stage dependency for agent
   * Any pipeline logic building the agent

5. Frontend build should simply copy from public/ to dist/

6. Rule:
   Rebuild and replace JAR ONLY when local_agent/java-agent code changes

Goal:

* Reduce pipeline complexity
* Avoid Artifactory/npm dependency issues
* Ensure stable and predictable builds
* Keep solution enterprise-friendly
