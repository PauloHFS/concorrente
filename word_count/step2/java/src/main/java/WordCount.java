import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

class FileCount implements Runnable {
    private String path;
    private AtomicInteger countRef;

    public FileCount(String filePath, AtomicInteger countref) {
        this.path = filePath;
        this.countRef = countref;
    }

    @Override
    public void run() {
       this.wcFile(this.path); 
    }

    public void wc(String fileContent) {
        String[] words = fileContent.split("\\s+");
        this.countRef.addAndGet(words.length);
    }
  
    public void wcFile(String filePath) {
        try {
            BufferedReader reader = new BufferedReader(new FileReader(filePath));
            StringBuilder fileContent = new StringBuilder();
            String line;
  
            while ((line = reader.readLine()) != null) {
                fileContent.append(line).append("\n");
            }
 
            reader.close();
            wc(fileContent.toString());
  
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

}

public class WordCount {
    
    // Calculate the number of words in the files stored under the directory name
    // available at argv[1].
    //
    // Assume a depth 3 hierarchy:
    //   - Level 1: root
    //   - Level 2: subdirectories
    //   - Level 3: files
    //
    // root
    // ├── subdir 1
    // │     ├── file
    // │     ├── ...
    // │     └── file
    // ├── subdir 2
    // │     ├── file
    // │     ├── ...
    // │     └── file
    // ├── ...
    // └── subdir N
    // │     ├── file
    // │     ├── ...
    // │     └── file
    public static void main(String[] args) {
        System.out.println("Running with " + Runtime.getRuntime().availableProcessors() + " cpus");
        if (args.length != 1) {
            System.err.println("Usage: java WordCount <root_directory>");
            System.exit(1);
        }

        String rootPath = args[0];
        File rootDir = new File(rootPath);
        File[] subdirs = rootDir.listFiles();
        AtomicInteger count = new AtomicInteger();
        
        ExecutorService executor = Executors.newFixedThreadPool(Runtime.getRuntime().availableProcessors());
       
        if (subdirs != null) {
            for (File subdir : subdirs) {
                if (subdir.isDirectory()) {
                    String dirPath = rootPath + "/" + subdir.getName();
                    File dir = new File(dirPath);
                    File[] files = dir.listFiles();
   
                    if (files != null) { 
                        for (File file : files) {
                            if (file.isFile()) {
                                Runnable worker = new FileCount(file.getAbsolutePath(), count);
                                executor.execute(worker);
                            }
                        }
                    }                    
                }
            }
        }

        executor.shutdown();
        while (!executor.isTerminated()) {}

        System.out.println(count);
    }

}
