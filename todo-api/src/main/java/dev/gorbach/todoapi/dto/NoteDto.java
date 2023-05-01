package dev.gorbach.todoapi.dto;

import java.sql.Timestamp;
import java.util.UUID;

import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
public class NoteDto {
    public UUID id;
    public Timestamp creationTimestamp;
    public String content;
    public Boolean isCompleted;
}